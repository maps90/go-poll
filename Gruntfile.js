module.exports = function(grunt) {
	'use strict';

	var globalConfig = {
		src: './deploy',
		dest: './public'
	};

	grunt.initConfig({
		globalConfig: globalConfig,
		pkg: grunt.file.readJSON('package.json'),
		banner: '/*!\n' +
			' * Dimbot v<%= pkg.version %> (<%= pkg.homepage %>)\n' +
			' * Copyright 2015-<%= grunt.template.today("yyyy") %> <%= pkg.author %>\n' +
			' */\n',
		jqueryCheck: 'if (typeof jQuery === \'undefined\') { throw new Error(\'Bootstrap\\\'s JavaScript requires jQuery\') }\n\n',
		uglify: {
			options: {
				banner: '/*! <%= pkg.name %> - v<%= pkg.version %> - ' +
				'<%= grunt.template.today("yyyy-mm-dd") %> */',
				beautify : {
					beautify: false,
					ascii_only: true,
					quote_keys: true
				}
			},
			build: {
				files: [{
					expand: true,
					src: ['**/*.js', '!**/*.min.js', '!**/*.backup.js'],
					dest: '<%= globalConfig.dest %>/js/',
					cwd: '<%= globalConfig.src %>/app/javascript/',
					extDot: 'last',
					ext: '.min.js'
				}]
			}
		},
		sass: {
			dist: {
				options: {
					banner: '<%= banner %>',
					style: 'expanded',
					quite: false,
				},
				files: {
					"<%= globalConfig.src %>/unminified/css/bootstrap.css": "<%= globalConfig.src %>/app/stylesheets/bootstrap.scss",
					"<%= globalConfig.src %>/unminified/css/dimbot-production.css": "<%= globalConfig.src %>/app/stylesheets/dimbot-production.scss",
					"<%= globalConfig.src %>/unminified/css/library.css": "<%= globalConfig.src %>/app/stylesheets/library.scss",
				}
			}
		},
		cssmin: {
			minify: {
				expand: true,
				src: ['*.css', '!*.min.css'],
				dest: '<%= globalConfig.dest %>/css/',
				cwd: '<%= globalConfig.src %>/unminified/css/',
				extDot: 'last',
				ext: '.min.css'
			}
		},
		watch: {
			files: ['<%= globalConfig.src %>/app/stylesheets/**/*.scss'],
			tasks: ['dist-sass'],
		}
	});

	grunt.loadNpmTasks('grunt-contrib-jshint');
	grunt.loadNpmTasks('grunt-contrib-sass');
	grunt.loadNpmTasks('grunt-contrib-uglify');
	grunt.loadNpmTasks('grunt-contrib-cssmin');
	grunt.loadNpmTasks('grunt-contrib-watch');
	grunt.registerTask('default', ['uglify', 'sass', 'cssmin']);
	grunt.registerTask('dist-sass', ['sass','cssmin']);
	grunt.registerTask('dist-js', ['uglify']);
	grunt.registerTask('dist-watch', ['watch']);
}
