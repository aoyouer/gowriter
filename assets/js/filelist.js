// 用于在posts中获取文件列表
function getFileList(type, catetgory, page, size) {
    fetch(`/api/filelist?type=${type}&category=${catetgory}&page=${page}&size=${size}`)
        .then((resp) => {
            resp.json().then((data) => {
                console.log(data)
                let fileList = ""
                data.files.forEach((file) => {
                    console.log(file)
                    console.log(file.FilePath)
                    fileList += `<div class="mdui-panel-item file-item">
                                    <div class="mdui-panel-item-header">
                                        <div class="mdui-panel-item-title">${file.FileName}</div>
                                        <div class="mdui-panel-item-summary">${file.FileModTime}</div>
                                        <i class="mdui-panel-item-arrow mdui-icon material-icons">keyboard_arrow_down</i>
                                    </div>
                                    <div class="mdui-panel-item-body">
                                    <p>content</p>
                                    <div class="mdui-panel-item-actions">
                                        <button class="mdui-btn mdui-ripple" mdui-panel-item-close>取消</button>
                                        <button class="mdui-btn mdui-ripple" onclick='window.location.href="/editor?path=${file.FilePath}"'>编辑</button>
                                    </div>
                                    </div>
                                 </div>`
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