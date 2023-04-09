package lg

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg"
	"github.com/golang-module/carbon"
	"github.com/nguyenthenguyen/docx"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func OpenLetter(order lg.Order, templateFile lg.File) (letter lg.Letter, file lg.File, encryptFile lg.File, err error) {
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10)
	out, err := os.Create(global.GVA_CONFIG.Insurance.TempDir + fileName + ".docx")
	if err != nil {
		return
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)
	_, err = out.Write(templateFile.FileSteam)
	if err != nil {
		return
	}

	currentTime := time.Now()
	var elogNo string
	if order.Delay != nil {
		elogNo = *order.Delay.ElogNo
	} else if order.Letter != nil {
		elogNo = *order.Letter.ElogNo
	} else {
		elogNo = global.GVA_CONFIG.Insurance.ElogPrefix + carbon.Now(carbon.Shanghai).Layout("20060102") + GenerateMD5String([]byte(*order.OrderNo + time.Now().String()))[:8]
	}
	insuredName := *order.Project.TendereeName
	insureName := *order.Apply.InsureName
	projectNo := *order.Project.ProjectNo
	projectName := *order.Project.ProjectName
	tenderDeposit := *order.Project.TenderDeposit
	tenderDepositString := fmt.Sprintf("%.2f", tenderDeposit)
	tenderDepositStringFixed, _ := strconv.ParseFloat(tenderDepositString, 64)
	tenderDepositCNY := ConvertNumToCny(tenderDepositStringFixed)
	insuranceName := global.GVA_CONFIG.Insurance.Name
	insuranceAddress := global.GVA_CONFIG.Insurance.Address
	insuranceZipCode := global.GVA_CONFIG.Insurance.ZipCode
	insuranceTel := global.GVA_CONFIG.Insurance.Tel
	year, month, day := currentTime.Date()
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tenderEndDate, _ := time.ParseInLocation("2006-01-02 15:04:05", *order.Project.ProjectOpenTime, loc)
	insureEndDate := tenderEndDate.AddDate(0, 0, int(*order.Project.ProjectDay))
	insureEndDateYear, insureEndDateMonth, insureEndDateDay := insureEndDate.Date()
	insureDay := *order.Project.ProjectDay
	insuredAddress := *order.Project.TendereeAddress
	var projectCategory string
	if order.Project.ProjectCategory != nil {
		projectCategory = *order.Project.ProjectCategory
	} else {
		projectCategory = ""
	}
	projectOpenDate, _ := time.ParseInLocation("2006-01-02 15:04:05", *order.Project.ProjectOpenTime, loc)
	projectOpenDateYear, projectOpenDateMonth, projectOpenDateDay := projectOpenDate.Date()
	projectPublishDate, _ := time.ParseInLocation("2006-01-02 15:04:05", *order.Project.ProjectPublishTime, loc)
	publishDateYear, publishDateMonth, publishDateDay := projectPublishDate.Date()

	insuranceCreditCode := global.GVA_CONFIG.Insurance.CreditCode
	urlString := GenerateMD5String([]byte(time.Now().String()))
	elogOutDate := currentTime.Format("2006-01-02 15:04:05")
	elogUrl := urlString[:16]
	elogEncryptUrl := urlString[16:]
	insureStartDateString := currentTime.Format("2006-01-02 15:04:05")
	insureEndDateString := insureEndDate.Format("2006-01-02 15:04:05")
	validateCode := urlString[12:20]

	var validateUrl string
	validateUrl = global.GVA_CONFIG.Insurance.HostDomain + "/elogValidate?elogNo=" + elogNo + "&validateCode=" + validateCode
	insureAddress := *order.Apply.InsureAddress

	letterReader, _ := docx.ReadDocxFile(global.GVA_CONFIG.Insurance.TempDir + fileName + ".docx")
	letterDocx := letterReader.Editable()
	_ = letterDocx.Replace("{elogNo}", elogNo, -1)
	_ = letterDocx.Replace("{insuredName}", insuredName, -1)
	_ = letterDocx.Replace("{insureName}", insureName, -1)
	_ = letterDocx.Replace("{projectNo}", projectNo, -1)
	_ = letterDocx.Replace("{projectName}", projectName, -1)
	_ = letterDocx.Replace("{tenderDeposit}", tenderDepositString, -1)
	_ = letterDocx.Replace("{tenderDepositCNY}", tenderDepositCNY, -1)
	_ = letterDocx.Replace("{insuranceName}", insuranceName, -1)
	_ = letterDocx.Replace("{insuranceAddress}", insuranceAddress, -1)
	_ = letterDocx.Replace("{insuranceZipCode}", insuranceZipCode, -1)
	_ = letterDocx.Replace("{insuranceTel}", insuranceTel, -1)
	_ = letterDocx.Replace("{year}", strconv.Itoa(year), -1)
	_ = letterDocx.Replace("{month}", strconv.Itoa(int(month)), -1)
	_ = letterDocx.Replace("{day}", strconv.Itoa(day), -1)
	_ = letterDocx.Replace("{insureStartDateYear}", strconv.Itoa(year), -1)
	_ = letterDocx.Replace("{insureStartDateMonth}", strconv.Itoa(int(month)), -1)
	_ = letterDocx.Replace("{insureStartDateDay}", strconv.Itoa(day), -1)
	_ = letterDocx.Replace("{insureEndDateYear}", strconv.Itoa(insureEndDateYear), -1)
	_ = letterDocx.Replace("{insureEndDateMonth}", strconv.Itoa(int(insureEndDateMonth)), -1)
	_ = letterDocx.Replace("{insureEndDateDay}", strconv.Itoa(insureEndDateDay), -1)
	_ = letterDocx.Replace("{insureDay}", strconv.FormatInt(insureDay, 10), -1)

	_ = letterDocx.Replace("{insuredAddress}", insuredAddress, -1)
	_ = letterDocx.Replace("{projectCategory}", projectCategory, -1)
	_ = letterDocx.Replace("{projectOpenDateYear}", strconv.Itoa(projectOpenDateYear), -1)
	_ = letterDocx.Replace("{projectOpenDateMonth}", strconv.Itoa(int(projectOpenDateMonth)), -1)
	_ = letterDocx.Replace("{projectOpenDateDay}", strconv.Itoa(projectOpenDateDay), -1)
	_ = letterDocx.Replace("{publishDateYear}", strconv.Itoa(publishDateYear), -1)
	_ = letterDocx.Replace("{publishDateMonth}", strconv.Itoa(int(publishDateMonth)), -1)
	_ = letterDocx.Replace("{publishDateDay}", strconv.Itoa(publishDateDay), -1)
	_ = letterDocx.Replace("{validateUrl}", validateUrl, -1)
	_ = letterDocx.Replace("{insureAddress}", insureAddress, -1)

	imageIndex := letterDocx.ImagesLen()
	if imageIndex > 0 {
		err = CreateQrCodeBs64WithLogo(validateUrl, global.GVA_CONFIG.Insurance.LogoFile, global.GVA_CONFIG.Insurance.TempDir+fileName+".png", 512)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + fileName + ".png")
		}()
		_ = letterDocx.ReplaceImage("word/media/image"+strconv.Itoa(imageIndex)+".png", global.GVA_CONFIG.Insurance.TempDir+fileName+".png")
	}

	_ = letterDocx.WriteToFile(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + ".docx")

	letterEncryptReader, _ := docx.ReadDocxFile(global.GVA_CONFIG.Insurance.TempDir + fileName + ".docx")
	letterEncryptDocx := letterEncryptReader.Editable()
	_ = letterEncryptDocx.Replace("{elogNo}", "**********", -1)
	_ = letterEncryptDocx.Replace("{insuredName}", "**********", -1)
	_ = letterEncryptDocx.Replace("{insureName}", "**********", -1)
	_ = letterEncryptDocx.Replace("{projectNo}", "**********", -1)
	_ = letterEncryptDocx.Replace("{projectName}", "**********", -1)
	_ = letterEncryptDocx.Replace("{tenderDeposit}", "**********", -1)
	_ = letterEncryptDocx.Replace("{tenderDepositCNY}", "**********", -1)
	_ = letterEncryptDocx.Replace("{insuranceName}", insuranceName, -1)
	_ = letterEncryptDocx.Replace("{insuranceAddress}", insuranceAddress, -1)
	_ = letterEncryptDocx.Replace("{insuranceZipCode}", insuranceZipCode, -1)
	_ = letterEncryptDocx.Replace("{insuranceTel}", insuranceTel, -1)
	_ = letterEncryptDocx.Replace("{year}", strconv.Itoa(year), -1)
	_ = letterEncryptDocx.Replace("{month}", strconv.Itoa(int(month)), -1)
	_ = letterEncryptDocx.Replace("{day}", strconv.Itoa(day), -1)
	_ = letterEncryptDocx.Replace("{insureStartDateYear}", strconv.Itoa(year), -1)
	_ = letterEncryptDocx.Replace("{insureStartDateMonth}", strconv.Itoa(int(month)), -1)
	_ = letterEncryptDocx.Replace("{insureStartDateDay}", strconv.Itoa(day), -1)
	_ = letterEncryptDocx.Replace("{insureEndDateYear}", strconv.Itoa(insureEndDateYear), -1)
	_ = letterEncryptDocx.Replace("{insureEndDateMonth}", strconv.Itoa(int(insureEndDateMonth)), -1)
	_ = letterEncryptDocx.Replace("{insureEndDateDay}", strconv.Itoa(insureEndDateDay), -1)
	_ = letterEncryptDocx.Replace("{insureDay}", strconv.FormatInt(insureDay, 10), -1)

	_ = letterEncryptDocx.Replace("{insuredAddress}", "**********", -1)
	_ = letterEncryptDocx.Replace("{projectCategory}", projectCategory, -1)
	_ = letterEncryptDocx.Replace("{projectOpenDateYear}", strconv.Itoa(projectOpenDateYear), -1)
	_ = letterEncryptDocx.Replace("{projectOpenDateMonth}", strconv.Itoa(int(projectOpenDateMonth)), -1)
	_ = letterEncryptDocx.Replace("{projectOpenDateDay}", strconv.Itoa(projectOpenDateDay), -1)
	_ = letterEncryptDocx.Replace("{publishDateYear}", strconv.Itoa(publishDateYear), -1)
	_ = letterEncryptDocx.Replace("{publishDateMonth}", strconv.Itoa(int(publishDateMonth)), -1)
	_ = letterEncryptDocx.Replace("{publishDateDay}", strconv.Itoa(publishDateDay), -1)
	_ = letterEncryptDocx.Replace("{validateUrl}", "**********", -1)
	_ = letterEncryptDocx.Replace("{insureAddress}", "**********", -1)
	_ = letterEncryptDocx.WriteToFile(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + "encrypt.docx")

	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + fileName + ".docx")

	err = exec.Command("libreoffice", "--headless", "--convert-to", "pdf", global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+".docx", "--outdir", global.GVA_CONFIG.Insurance.TempDir).Run()
	if err != nil {
		return
	}
	err = exec.Command("libreoffice", "--headless", "--convert-to", "pdf", global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+"encrypt.docx", "--outdir", global.GVA_CONFIG.Insurance.TempDir).Run()
	if err != nil {
		return
	}

	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + ".docx")
	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + "encrypt.docx")

	err = exec.Command("java", "-jar", global.GVA_CONFIG.Insurance.SignProgram, global.GVA_CONFIG.Insurance.KeyFile, global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+".pdf", global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+"Signed.pdf", global.GVA_CONFIG.Insurance.StampFile, global.GVA_CONFIG.Insurance.LegalFile).Run()
	if err != nil {
		return
	}
	err = exec.Command("java", "-jar", global.GVA_CONFIG.Insurance.SignProgram, global.GVA_CONFIG.Insurance.KeyFile, global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+"encrypt.pdf", global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+"encryptSigned.pdf", global.GVA_CONFIG.Insurance.StampFile, global.GVA_CONFIG.Insurance.LegalFile).Run()
	if err != nil {
		return
	}

	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + ".pdf")
	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + "encrypt.pdf")

	letterFileName := "letter" + fileName + "Signed.pdf"
	encryptLetterFileName := "letter" + fileName + "encryptSigned.pdf"
	letterFile, err := os.Open(global.GVA_CONFIG.Insurance.TempDir + letterFileName)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(letterFile)
	letterFileContent, err := io.ReadAll(letterFile)
	if err != nil {
		return
	}
	encryptLetterFile, err := os.Open(global.GVA_CONFIG.Insurance.TempDir + encryptLetterFileName)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(encryptLetterFile)
	encryptLetterFileContent, err := io.ReadAll(encryptLetterFile)
	if err != nil {
		return
	}

	file = lg.File{
		FileName:  &letterFileName,
		FileSteam: letterFileContent,
	}
	encryptFile = lg.File{
		FileName:  &encryptLetterFileName,
		FileSteam: encryptLetterFileContent,
	}
	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + letterFileName)
	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + encryptLetterFileName)
	letter = lg.Letter{
		OrderID:             &order.ID,
		OrderNo:             order.OrderNo,
		ElogNo:              &elogNo,
		InsuranceName:       &insuranceName,
		InsuranceCreditCode: &insuranceCreditCode,
		ElogOutDate:         &elogOutDate,
		ElogUrl:             &elogUrl,
		ElogEncryptUrl:      &elogEncryptUrl,
		TenderDeposit:       order.Project.TenderDeposit,
		InsureStartDate:     &insureStartDateString,
		InsureEndDate:       &insureEndDateString,
		InsureDay:           order.Project.ProjectDay,
		ValidateCode:        &validateCode,
	}
	return
}
