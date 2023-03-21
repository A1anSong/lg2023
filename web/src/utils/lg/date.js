export const date = (date) => {
  if (date === '') {
    return ''
  }
  const dateUTC = new Date(date)
  return dateUTC.getFullYear() +
        '-' + ('0' + (dateUTC.getMonth() + 1)).slice(-2) +
        '-' + ('0' + dateUTC.getDate()).slice(-2) +
        ' ' + dateUTC.toLocaleTimeString('chinese', { hour12: false })
}
