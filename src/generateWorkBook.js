const Excel = require("exceljs");
const moment = require("moment");
const { slice1 } = require('./wattHour');
const {
  METER_NAME,
  BORDERS,
  EXCLUDED_METERS,
  EXCLUDED_METER_CONSUMPTION,
  EXCEL_FONTSTYLES,
  EXCEL_HORIZONTALLEFT,
  EXCEL_HEADING,
  EXCEL_ALIGNMENTSTYLES,
  WORKSHEET_NAMES,
  TITLE_FONT,
  SUBTITLE_FONT,
  MERGE_CELL_HEADERALIGNMENT,
} = require("./constants");
const { LOGO } = require("./logobase64");

generateWorkBook = (res, workBookDatas) => {
  var workbook = new Excel.Workbook();

  const fromDate = moment(workBookDatas[0].from).format("DD-MM-YYYY HH:mm:ss");
  const toDate = moment(workBookDatas[0].to).format("DD-MM-YYYY HH:mm:ss");
  const logoBase64Image = LOGO;
  const logo = workbook.addImage({
    base64: logoBase64Image,
    extension: "png",
  });

  /** Meter Reading Sheet */
  var meterSheet = workbook.addWorksheet(WORKSHEET_NAMES[0]);

  //Main Header
  meterSheet.mergeCells("A1:E3");

  meterSheet.addImage(logo, "A1:A3");

  const meterCell = meterSheet.getCell("B2");

  meterCell.font = EXCEL_FONTSTYLES;
  meterCell.alignment = EXCEL_ALIGNMENTSTYLES;
  meterCell.value = {
    richText: [
      {
        font: TITLE_FONT,
        text: `${EXCEL_HEADING}\n`,
      },
      {
        font: SUBTITLE_FONT,
        text: `From: ${fromDate} To: ${toDate}`,
      },
    ],
  };
  meterCell.alignment = MERGE_CELL_HEADERALIGNMENT;
  meterCell.border = BORDERS;


  //OtherReadings
  var meterHeadingRow = meterSheet.addRow();
  var meterColumns = workBookDatas[0].column_name;
  var meterDatas = workBookDatas[0].sheet_data;
  var meterNames = Object.values(METER_NAME);

  meterSheet.getRow(4).font = EXCEL_FONTSTYLES;

  //setting header
  for (let i = 0; i < meterColumns.length; i++) {
    let currentColumnWidth = meterColumns[i].width;
    meterSheet.getColumn(i + 1).width = currentColumnWidth;
    let cell = meterHeadingRow.getCell(i + 1);
    cell.value = meterColumns[i].header;
    cell.border = BORDERS;
  }
  

  //adding rowvalues
  for (let i = 0; i < meterNames.length; i++) {
    const meterConsumption = meterDatas.find(
      (meterData) => meterData.name === meterNames[i]
    );
    var meterRow = meterSheet.addRow();
    for (let j = 0; j < 3; j++) {
      meterRow.getCell(1).value = i + 1;
      meterRow.getCell(2).value = meterNames[i];
      meterRow.getCell(3).value = meterConsumption === undefined ? 0 : meterConsumption.start;
      meterRow.getCell(4).value = meterConsumption === undefined ? 0 : meterConsumption.end;
      meterRow.getCell(5).value =
        meterConsumption === undefined ? 0 : meterConsumption.consumption;
      meterRow.getCell(1).alignment = EXCEL_HORIZONTALLEFT;
      meterRow.getCell(1).border = BORDERS;
      meterRow.getCell(2).border = BORDERS;
      meterRow.getCell(3).border = BORDERS;
      meterRow.getCell(4).border = BORDERS;
      meterRow.getCell(5).border = BORDERS;
    }
  }

  /** Overall Excel Sheet */

  var worksheet = workbook.addWorksheet(WORKSHEET_NAMES[1]);
  var serialNumber = 0;

  // Main Header

  worksheet.mergeCells("A1:E3");
  worksheet.addImage(logo, "A1:A3");

  const customCell = worksheet.getCell("A2");

  customCell.font = EXCEL_FONTSTYLES;
  customCell.alignment = EXCEL_ALIGNMENTSTYLES;
  // customCell.value = EXCEL_HEADING;
  customCell.value = {
    richText: [
      {
        font: TITLE_FONT,
        text: `${EXCEL_HEADING} \n`,
      },
      {
        font: SUBTITLE_FONT,
        text: `From: ${fromDate} To: ${toDate}`,
      },
    ],
  };
  customCell.alignment = MERGE_CELL_HEADERALIGNMENT;
  customCell.border = BORDERS;

  //Main Incoming Data
  worksheet.mergeCells("A4:C4");
  const mainIncomingCell = worksheet.getCell("A4");
  mainIncomingCell.alignment = EXCEL_ALIGNMENTSTYLES;
  mainIncomingCell.value = METER_NAME["Energymeter 2"];
  
  worksheet.mergeCells("D4:E4");
  const mainIncomingValue = worksheet.getCell("D4");
  const mainIncomingFilteredValue = workBookDatas[0].sheet_data.find(
    (element) => element.name == "MainIncoming"
  );
  mainIncomingValue.value =
    mainIncomingFilteredValue == undefined
      ? 0
      : mainIncomingFilteredValue.consumption;

  mainIncomingCell.font = EXCEL_FONTSTYLES;
  mainIncomingValue.font = EXCEL_FONTSTYLES;
  mainIncomingCell.border = BORDERS;
  mainIncomingValue.border = BORDERS;

  //Generator Data
  worksheet.mergeCells("A5:C5");
  const generatorCell = worksheet.getCell("A5");
  generatorCell.alignment = EXCEL_ALIGNMENTSTYLES;
  generatorCell.value = METER_NAME["Energymeter 3"];
  generatorCell.font = EXCEL_FONTSTYLES;
  
  worksheet.mergeCells("D5:E5");
  const generatorValue = worksheet.getCell("D5");
  const generatorFilteredValue = workBookDatas[0].sheet_data.find(
    (element) => element.name == "Generator"
  );
  generatorValue.value =
    generatorFilteredValue == undefined
      ? 0
      : generatorFilteredValue.consumption;

  generatorValue.font = EXCEL_FONTSTYLES;
  generatorCell.border = BORDERS;
  generatorValue.border = BORDERS;

  //Solar Data
  worksheet.mergeCells("A6:C6");
  const solarCell = worksheet.getCell("A6");
  solarCell.alignment = EXCEL_ALIGNMENTSTYLES;
  solarCell.value = METER_NAME["Energymeter 29"];
  
  worksheet.mergeCells("D6:E6");
  const solarValue = worksheet.getCell("D6");
  const solarFilteredValue = workBookDatas[0].sheet_data.find(
    (element) => element.name == "Solar"
  );
  solarValue.value =
    solarFilteredValue == undefined ? 0 : solarFilteredValue.consumption;

  solarCell.font = EXCEL_FONTSTYLES;
  solarValue.font = EXCEL_FONTSTYLES;
  solarCell.border = BORDERS;
  solarValue.border = BORDERS;

  //Total Incoming
  worksheet.mergeCells("A7:C7");
  const totalCell = worksheet.getCell("A7");
  totalCell.alignment = EXCEL_ALIGNMENTSTYLES;
  totalCell.value = "Total Incoming";
  
  worksheet.mergeCells("D7:E7");
  const totalIncomingValue = worksheet.getCell("D7");
  const totalSum =
    mainIncomingValue.value + generatorValue.value + solarValue.value;
  totalIncomingValue.value = totalSum;

  totalCell.font = EXCEL_FONTSTYLES;
  totalIncomingValue.font = EXCEL_FONTSTYLES;
  totalCell.border = BORDERS;
  totalIncomingValue.border = BORDERS;

  //OtherReadings
  var headerRow = worksheet.addRow();
  var columns = workBookDatas[0].column_name;
  var allItems = workBookDatas[0].sheet_data;
  worksheet.getRow(8).font = EXCEL_FONTSTYLES;


  //setting header
  for (let i = 0; i < columns.length; i++) {
    let currentColumnWidth = columns[i].width;
    worksheet.getColumn(i + 1).width = currentColumnWidth;
    let cell = headerRow.getCell(i + 1);
    cell.value = columns[i].header;
    cell.border = BORDERS;
  }

  //adding rowvalues
  for (let i = 0; i < meterNames.length; i++) {
    if (!EXCLUDED_METERS.includes(meterNames[i])) {
      const energyConsumption = allItems.find(
        (allItem) => allItem.name === meterNames[i]
      );
      var dataRow = worksheet.addRow();
      serialNumber = serialNumber + 1;
      for (let j = 0; j < 3; j++) {
        dataRow.getCell(1).value = serialNumber;
        dataRow.getCell(2).value = meterNames[i];
        dataRow.getCell(3).value =
          energyConsumption === undefined ? 0 : energyConsumption.start;
        dataRow.getCell(4).value =
          energyConsumption === undefined ? 0 : energyConsumption.end;
        dataRow.getCell(5).value =
          energyConsumption === undefined ? 0 : energyConsumption.consumption;
        dataRow.getCell(1).alignment = EXCEL_HORIZONTALLEFT;
        dataRow.getCell(1).border = BORDERS;
        dataRow.getCell(2).border = BORDERS;
        dataRow.getCell(3).border = BORDERS;
        dataRow.getCell(4).border = BORDERS;
        dataRow.getCell(5).border = BORDERS;
      }
    }
  }

  //setting footer
  const rowCount = worksheet.rowCount;
  const footerTextCell = `A${rowCount + 1}:D${rowCount + 2}`;
  const footerTextValue = `E${rowCount + 1}:E${rowCount + 2}`;
  var overallConsumption = 0;
  for (let i = 0; i < allItems.length; i++) {
    if (!EXCLUDED_METER_CONSUMPTION.includes(allItems[i].name)) {
      overallConsumption += allItems[i].consumption;
    }
  }

  worksheet.mergeCells(footerTextCell);
  worksheet.getRow(1).font = EXCEL_FONTSTYLES;
  worksheet.getCell(`A${rowCount + 1}`).alignment = EXCEL_ALIGNMENTSTYLES;
  worksheet.getCell(`A${rowCount + 1}`).value = "Overall Consumption";
  worksheet.getCell(`A${rowCount + 1}`).font = EXCEL_FONTSTYLES;
  worksheet.mergeCells(footerTextValue);
  worksheet.getCell(`C${rowCount + 1}`).alignment = {
    vertical: "middle",
    horizontal: "center",
  };
  worksheet.getCell(`E${rowCount + 1}`).value = overallConsumption;
  worksheet.getCell(`E${rowCount + 1}`).font = EXCEL_FONTSTYLES;
  worksheet.getCell(footerTextCell).border = BORDERS;
  worksheet.getCell(`E${rowCount + 1}`).border = BORDERS;

  /** Creating workbook */

  res.setHeader(
    "Content-Type",
    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
  );
  res.setHeader(
    "Content-Disposition",
    "attachment; filename=" + `Energymeter-${Date.now()}.xlsx`
  );

  return workbook.xlsx.write(res).then(function (result) {
    console.log("excel downloaded successfully");
    res.status(200).end();
  });
};

module.exports = {
  generateWorkBook,
};
