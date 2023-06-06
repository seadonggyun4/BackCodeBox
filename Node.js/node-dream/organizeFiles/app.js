const fs = require('fs').promises

let files = []


async function  organizeFiles () {
    // [ 현재 경로에 있는 파일이름, 경로이름 받아오기 ]
    files = await fs.readdir('./')
        .then((res) => {
            return  res
        })
        .catch((e) => {
            console.log(e)
            return []
        })


    // [ 폴더들 생성 ]
    const forderName = ['duplicated', 'video', 'captured']

    makeForders(forderName)

    // [ 파일 확장자 검사 및 정렬 ]
    await checkType(files)
}


function makeForders (forderName) {
    forderName.forEach(async (name) => {
        await fs.mkdir(name)
            .catch(console.error)
    })
}


function checkType (files){
    files.forEach( async (file) => {
        const type = file.split(".").pop()

        if(type === 'jpg'){
            await fs.copyFile(`./${file}`, `./duplicated/${file}`)
                .catch(console.error)
        }


        if(type === 'mp4'){
            await fs.copyFile(`./${file}`, `./video/${file}`)
                .catch(console.error)

            await fs.unlink(`./${file}`)
                .catch(console.error)
        }


        if(type !== 'jpg' && type !== 'mp4' && file !== 'app.js'){
            await fs.copyFile(`./${file}`, `./captured/${file}`)
                .catch(console.error)

            await fs.unlink(`./${file}`)
                .catch(console.error)
        }
    })
}




organizeFiles()



