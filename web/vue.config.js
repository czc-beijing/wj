module.exports = {
    devServer: {
    },
    lintOnSave: false,
    publicPath: '/',
    configureWebpack: {
        module:{
            rules:[{
                test: /\.mjs$/,
                include: /node_modules/,
                type: 'javascript/auto'
            }]
        }
    }
}