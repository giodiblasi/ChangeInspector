const initDetail = () => {
    const getData = (url) => {
        var xmlhttp = new XMLHttpRequest();
        var response = new Promise(resolve=>{
            xmlhttp.onreadystatechange = function() {
                if (this.readyState == 4 && this.status == 200) {
                    resolve(JSON.parse(this.responseText));
                }
            }
        });
        xmlhttp.open("GET", url, true);
        xmlhttp.send();
        return response;
    }

    return{
        update: async (fileName) => console.log(await getData(`/detail/${fileName.replace(/\//g,'$')}`))
    }
}