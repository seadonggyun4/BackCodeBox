const http = require('http')
const fs = require('fs').promises


http.createServer(async (req, res) => {    
    try{
        if(req.url === '/'){// 주소가 / 일때
            const fdata = await fs.readFile('./index.html')
            res.writeHead(200, {'Content-Type' : 'text/html; charset=utf-8'})
            return res.end(fdata)
        }
        else if(req.url === '/about'){// 주소가 /about 일때
            const fdata = await fs.readFile('./about.html')
            res.writeHead(200, {'Content-Type' : 'text/html; charset=utf-8'})
            return res.end(fdata)
        }
        
        try{
            const fdata = await fs.readFile(`.${req.url}`)
            return res.end(fdata)
        }
        catch(err){
            res.writeHead(404)
            res.end('Page Not Found')
        }
        
    }
    catch(err){
        console.error(err)
    }
}).listen(8080, () => {
    console.log('8080번 포트 실행')
})