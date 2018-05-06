"use strict"

var moment = require('moment-timezone');

module.exports = (context, callback) => {
    var now = Date.now();
    var nowD = new Date(now);
    var thisMoment = moment(nowD);

    if (context == "list" || context == "ls") {
        var names = moment.tz.names();
        names.forEach(function(value) {
            console.log(value);
        });
    }
    else {
        console.log(thisMoment.tz(context).format());
    }
}
