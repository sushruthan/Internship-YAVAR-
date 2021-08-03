const { InfluxDB, flux } = require("@influxdata/influxdb-client");
const { url, token, org, bucket, timeout } = require("../env");
const { METER_NAME } = require('./constants');


getWattHourReceived = async (startDuration, endDuration) => {
  const queryApi = new InfluxDB({ url, token, timeout }).getQueryApi(org);

  //querying consumtion data
  const fluxQuerySUM = flux`from(bucket:"${bucket}") |> range(start: ${startDuration}, stop: ${endDuration}) 
  |> filter(fn: (r) => r._measurement == "modbus") 
  |> filter(fn: (r) => r["_field"] == "Wh Received") 
  |> map(fn: (r) => ({ r with _value: r._value / 1000.0}))  
  |> aggregateWindow(every: 1h, fn: last, createEmpty: false) 
  |> difference() |> sum()`;
  const fluxQuerySTART = flux`from(bucket:"${bucket}") |> range(start: ${startDuration}, stop: ${endDuration})   
  |> filter(fn: (r) => r._measurement == "modbus") 
  |> filter(fn: (r) => r["_field"] == "Wh Received") 
  |> map(fn: (r) => ({ r with _value: r._value / 1000.0}))  
  |> aggregateWindow(every: 1h, fn: first, createEmpty: false) 
  |> first()`;
  const fluxQueryEND = flux`from(bucket:"${bucket}") |> range(start: ${startDuration}, stop: ${endDuration})   
  |> filter(fn: (r) => r._measurement == "modbus") 
  |> filter(fn: (r) => r["_field"] == "Wh Received") 
  |> map(fn: (r) => ({ r with _value: r._value / 1000.0}))  
  |> aggregateWindow(every: 1h, fn: first, createEmpty: false) 
  |> last()`;


  const wattHourJSON = [];
  const wattHourJSONstart = [];
  const wattHourJSONend = [];
  var filteredWH = [];
  await queryApi
    .collectRows(fluxQuerySUM)
    .then((datas) => {
      datas.forEach((data) => {
        const meterName = METER_NAME[data.name];
        const value = data._value === null ? 0.00 : data._value;
        wattHourJSON.push({
          time: data._time,
          measurement: data._measurement,
          name: meterName,
          consumption: parseFloat(value.toFixed(2)),
          host: data.host,
          field: data._field,
        });
      });
      console.log("\nCollected wh received", wattHourJSON.length);
    })
    .catch((error) => {
      console.error(error);
      console.log("\nError in wh received row");
    });
    await queryApi
    .collectRows(fluxQuerySTART)
    .then((datas) => {
      datas.forEach((data) => {
        const meterName = METER_NAME[data.name];
        const value = data._value === null ? 0.00 : data._value;
        wattHourJSONstart.push({
          time: data._time,
          measurement: data._measurement,
          name: meterName,
          start: parseFloat(value.toFixed(2)),
          host: data.host,
          field: data._field,
        });
      });

      console.log("\nCollected wh start", wattHourJSONstart.length);
    })
    .catch((error) => {
      console.error(error);
      console.log("\nError in wh received row");
    });

    await queryApi
    .collectRows(fluxQueryEND)
    .then((datas) => {
      datas.forEach((data) => {
        const meterName = METER_NAME[data.name];
        const value = data._value === null ? 0.00 : data._value;
        wattHourJSONend.push({
          time: data._time,
          measurement: data._measurement,
          name: meterName,
          end: parseFloat(value.toFixed(2)),
          host: data.host,
          field: data._field,
        });
      });
      console.log("\nCollected wh end", wattHourJSONend.length);
    })
    .catch((error) => {
      console.error(error);
      console.log("\nError in wh received row");
    });

//merging array of object ID

    [wattHourJSON, wattHourJSONstart, wattHourJSONend].forEach(function (a) {
      a.forEach(function (b) {
          if (!this[b.name]) {
              this[b.name] = {};
              filteredWH.push(this[b.name]);
          }
          Object.keys(b).forEach(function (k) {
              this[b.name][k] = b[k];
          }, this);
      }, this);
  }, Object.create(null));
  return filteredWH;
};

module.exports = {
    getWattHourReceived,

};
