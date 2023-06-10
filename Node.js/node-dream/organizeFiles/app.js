const fs = require('fs').promises
const FS = require('fs')

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


async function makeForders (forderName) {
    for(const name of forderName){
        if(FS.existsSync(__dirname + `/${name}`)){
            return
        }

        await fs.mkdir(name)
            .catch(console.error)
    }
}


async function checkType (files){
    for(const file of files){
        const type = file.split(".").pop()

        if(type === 'jpg'){
            if(FS.existsSync(__dirname + `/duplicated/${file}`)){
                return
            }

            await fs.copyFile(`./${file}`, `./duplicated/${file}`)
                .catch(console.error)
        }


        if(type === 'mp4' || type === 'mov'){
            if(FS.existsSync(__dirname + `/video/${file}`)){
                return
            }

            await fs.copyFile(`./${file}`, `./video/${file}`)
                .catch(console.error)

            await fs.unlink(`./${file}`)
                .catch(console.error)
        }


        if(type !== 'jpg' && type !== 'mp4' && file !== 'app.js'){
            if(FS.existsSync(__dirname + `/captured/${file}`)){
                return
            }


            await fs.copyFile(`./${file}`, `./captured/${file}`)
                .catch(console.error)

            await fs.unlink(`./${file}`)
                .catch(console.error)
        }
    }
}




organizeFiles()


