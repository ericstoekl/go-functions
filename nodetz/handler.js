"use strict"

var moment = require('moment-timezone');

module.exports = (context, callback) => {
    var now = Date.now();
    var nowD = new Date(now);
    var thisMoment = moment(nowD);

    console.log(moment.tz.names());
    /*
    if (context == "list" || context == "ls") {
        moment.tz.names();
    }
    else {
        console.log(thisMoment.tz(context));
    }
    */
}
