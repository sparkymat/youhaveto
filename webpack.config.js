module.exports = {
    entry: "./js/index.js",
    output: {
        path: __dirname,
        filename: "public/js/youhaveto.js"
    },
    module: {
        loaders: [
            { test: /\.css$/, loader: "style!css" },
            { test: /\.js$/, exclude: /node_modules/, loader: "babel-loader", query: {presets: ["stage-1", "es2015", "react"], plugins: ["transform-class-properties"]} },
            { test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader:"url?limit=10000&mimetype=application/font-woff" },
      		{ test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader: "file" }
        ]
    }
};
