const fs = require('fs').promises

const handleFile = async () => {
    // [ wirte a file ]
    await fs.writeFile('./text.txt', 'hello Dream Coders! :) ')
        .catch(console.error)

    // [ append a file ]
    await fs.appendFile('./text.txt', '\n Yo Dream Coders! :) ')
        .catch(console.error)



    for(let i = 0; i < 5; i++){
        // [ make folder ]
        await fs.mkdir(`sub-folder${i}`)
            .catch(console.error)


        // [ copy a file ]
        await fs.copyFile('./text.txt', `./sub-folder${i}/text${i}.txt`)
            .catch(console.error)


        // [ read a file ]
        // - 일반적으로 파일을 읽을때는 buffer와 같은 형식으로 읽힌다
        // - readFile() 후미에 읽은 파일을 어떤 형식으로 인코딩 할 것인지 지정가능
        await fs.readFile(`./sub-folder${i}/text.txt`, 'utf8')
            .then((data) => console.log(data))
            .catch(console.error)
    }



    // [ 현재 경로에 있는 파일이름, 경로이름 받아오기 ]
    await fs.readdir('./')
        .then(console.log)
        .catch(console.error)
}


handleFile()

