import { Command } from 'commander';
const program = new Command();

program
    .name('fitogether')
    .version('0.1.0');

program.command('build-report')
    .description('Build cFit report for project')
    .option('--cfit-dir', 'cFit documents root directory')
    .option('--tests-results-dir', 'Test results root directory')
    .option('--output-dir', 'cFit report output root directory')
    .action((options) => {

    })

program.parse();