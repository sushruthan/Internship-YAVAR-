/** InfluxDB v2 URL localhost:3333/excelreport */
const url = 'http://localhost:8086'
/** InfluxDB authorization token */
const token = '9J2PHWZGzEacYjshCAP0OxAPx_fofocHtjemZa_ebKhfxvBPcsztm5Efy9jSX3wDO5VGwG3-jIqdXYH2AnbZow=='
/** Organization within InfluxDB  */
const org = 'demo'
/**InfluxDB bucket used in examples  */
const bucket = 'example-bucket'
// ONLY onboarding example
/**InfluxDB user  */
const username = 'Artemis'
/**InfluxDB password  */
const password = 'YJSKyH2Cgb9xfHN'

const timeout = 500000

module.exports = {
  url,
  token,
  org,
  bucket,
  username,
  password,
  timeout,
}
