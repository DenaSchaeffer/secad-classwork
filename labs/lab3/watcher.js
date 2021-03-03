const fs = require('fs'); //file system module in Node
fs.watch('target.txt', function() {//the callback
console.log("File 'target.txt' just changed!");
});
console.log("Lab 3-3.a.ii by Dena Schaeffer. Watching target.txt for changes...");