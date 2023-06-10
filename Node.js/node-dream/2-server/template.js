const http = require('http')
const ejs = require('ejs')

const name = 'DGSEO'
const courses = [{name: 'HTML'}, {name: 'CSS'}, {name: 'JS'} ,{name: 'NODE'} ,{name: 'VUE'} ,{name: 'REACT'}]

const server = http.createServer((req,  res) => {
    const URL = req.url
    res.setHeader('Content-Type', 'text/html')

    if(URL === '/'){
        ejs.renderFile('./template/home.ejs', {name})
            .then(data => res.end(data))
    }

    if(URL === '/courses'){
        ejs.renderFile('./template/courses.ejs', {name, courses})
            .then(data => res.end(data))
    }
    res.end()
})




server.listen(8888)