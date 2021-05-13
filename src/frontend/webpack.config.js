const path = require('path');

module.exports = {
	resolve: {
		alias: {
			'node_modules': path.join(__dirname, 'node_modules'),
			'icons': path.join(__dirname, 'assets/icons')
		},
		extensions: ['*', '.js', '.jsx'],
	},
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: ['babel-loader', 'source-map-loader'],
      },
      {
        test: /\.css$/i,
        use: ["style-loader", "css-loader"],
      },
			{
        test: /\.(png|svg|jpg|jpeg|gif)$/i,
        type: 'asset/resource',
      },
    ],
  },
  entry: {
    main: path.resolve(__dirname, './js/main.js'),
  },
  output: {
    filename: '[name].bundle.js',
    path: path.resolve(__dirname, 'build')
  },
};

