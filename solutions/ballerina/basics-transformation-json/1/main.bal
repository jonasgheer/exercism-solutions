import ballerina/io;

type FillUpEntry record {|
    int employeeId;
    int odometerReading;
    decimal gallons;
    decimal gasPrice;
|};

type EmployeeResult record {|
    readonly int employeeId;
    int gasFillUpCount;
    decimal totalFuelCost;
    decimal totalGallons;
    int totalMilesAccrued;
|};

function processEmployee(FillUpEntry[] entries) returns EmployeeResult {
    var odometerReadings = entries.map(e => e.odometerReading).sort("ascending");
    return {
        employeeId: entries[0].employeeId,
        gasFillUpCount: entries.length(),
        totalFuelCost: entries.reduce(function(decimal total, FillUpEntry entry) returns decimal => (entry.gallons * entry.gasPrice) + total, 0),
        totalGallons: entries.reduce(function(decimal total, FillUpEntry entry) returns decimal => entry.gallons + total, 0),
        totalMilesAccrued: odometerReadings.pop() - odometerReadings.shift()
    };
}

function groupBy(FillUpEntry[] entries) returns map<FillUpEntry[]> {
    map<FillUpEntry[]> groups = {};
    foreach var entry in entries {
        var employeeId = entry.employeeId.toString();
        if !groups.hasKey(employeeId) {
            groups[employeeId] = [entry];
            continue;
        }
        groups.get(employeeId).push(entry);
    }
    return groups;
}

function processFuelRecords(string inputFilePath, string outputFilePath) returns error? {
    json entriesJson = check io:fileReadJson(inputFilePath);
    FillUpEntry[] entries = check entriesJson.cloneWithType();

    var entriesById = groupBy(entries);

    var results = from var r in entriesById.map(processEmployee).toArray()
        order by r.employeeId
        select r;

    check io:fileWriteJson(outputFilePath, results.toJson());
}

