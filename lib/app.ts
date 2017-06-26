import yargs = require('yargs');

yargs.commandDir('./commands')
  .help()
  .argv;
