var webpack = require('webpack');

var commonsPlugin = new webpack.optimize.CommonsChunkPlugin('common.js');

module.exports = {
	entry: {
		'/bundle': './public/src/js/main.js',
		'../../meshScript/dist/savingMesh': './meshScript/src/savingMesh.js'
	},
	output: {
		path: __dirname + '/public/js',
		filename: '[name].js'
	},
	plugins: [
		new webpack.ProvidePlugin({ riot: 'riot' }),
		commonsPlugin
		// new webpack.optimize.UglifyJsPlugin()
	],
	module: {
		preLoaders: [
			{
				test: /\.tag$/,
				exclude: /node_modules/,
				loader: 'riotjs-loader',
				query: {
					type: ['babel', 'scss']
				}
			}
		],
		loader: [
			{
				test: /\.js$|\.tag$/,
				exclude: /node_modules/,
				loader: 'babel-loader',
			}
		]
	},
	resolve: {
		extensions: ['', '.js', '.tag']
	}
};
