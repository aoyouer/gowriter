// 用于在posts中获取文件列表
function getFileList(type, catetgory, page, size) {
    fetch(`/api/filelist?type=${type}&category=${catetgory}&page=${page}&size=${size}`)
        .then((resp) => {
            resp.json().then((data) => {
                console.log(data)
                let fileList = ""
                data.files.forEach((file) => {
                    console.log(file)
                    fileList += `<li class="mdui-list-item mdui-ripple"><div class="mdui-list-item-content">${file.FileName}</div></li>`
                })
                mdui.$("#filelist").append(fileList)
            })
        })
        .catch((error) => {
            console.log(error)
        })
}

function initFilelist() {
    getFileList('posts', '', 1, 10)
}

window.onload = function () {
    initFilelist()
}