module.exports = function(grunt) {
	"use strict";
	grunt.initConfig({
	beep: false,
	pkg: grunt.file.readJSON('package.json'),
	watch: {
		files: ['**/*.go'],
		tasks: ['gobuild']
		}
	});
	grunt.loadNpmTasks('grunt-contrib-watch');
	grunt.option('beep',false);
	grunt.registerTask('gobuild',"runs go build command",function(){
		var exec = require('child_process').exec;
		var cb = this.async();
		exec('go build', {cwd: '.'}, function(err, stdout, stderr) {
			console.log(stdout,stderr);
			cb();
		});
	});
};