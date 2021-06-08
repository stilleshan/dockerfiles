const gulp = require("gulp");
const buble = require("gulp-buble");
const uglify = require("gulp-uglify");
const rename = require("gulp-rename");

function build() {
  return gulp
    .src("./src/**/*.js")
    .pipe(buble())
    .pipe(gulp.dest("./dist"))
    .pipe(
      uglify({
        mangle: true,
        ie8: true
      })
    )
    .pipe(rename({ suffix: ".min" }))
    .pipe(gulp.dest("./dist"));
}

exports.build = build;
