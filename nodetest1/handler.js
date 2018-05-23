"use strict"

module.exports = (context, callback) => {
    console.log("Hello world");
    callback(undefined, {status: "done"});
}
