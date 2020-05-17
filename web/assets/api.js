const buildClient = ()=>{

    const apiListeners={}

    const subscribe = (event, callback) => {
        const listeners = apiListeners[event];
        console.log('subscription')
        if(!listeners){
            apiListeners[event]=[callback]
        }
        else{
            apiListeners[event]=[...listeners, callback]
        }
    }

    const notify = (event, payload) =>{
        return new Promise(resolve=>{
            setTimeout(function(){
                listeners = apiListeners[event]
                if(listeners){
                    listeners.forEach(clbk=>{console.log('calling');clbk(payload)})
                }
                resolve();
            },100);
        });
    }
    const getFilter = () => {
        var xmlhttp = new XMLHttpRequest();
        var response = new Promise(resolve => {
            xmlhttp.onreadystatechange = function () {
                if (this.readyState == 4 && this.status == 200) {
                    const response = JSON.parse(this.responseText);
                    resolve(response);
                }
            }
        });
        xmlhttp.open("GET", '/filter', true);
        xmlhttp.send(); 
        return response;
    }

    const removeFromFilter = (fileName) => {
        var xmlhttp = new XMLHttpRequest();
        var response = new Promise(resolve => {
            xmlhttp.onreadystatechange = function () {
                if (this.readyState == 4 && this.status == 200) {
                    notify("FILTER_ITEM_REMOVED", fileName);
                    resolve();
                }
            }
        });
        xmlhttp.open("DELETE", `/filter/${fileName.replace(/\//g, '$')}`, true);
        xmlhttp.send(); 
        return response;
    }

    const addToFilter = (fileName) => {
        var xmlhttp = new XMLHttpRequest();
        var response = new Promise(resolve => {
            xmlhttp.onreadystatechange = function () {
                if (this.readyState == 4 && this.status == 200) {
                    console.log("notification")
                    notify("FILTER_ITEM_ADDED", fileName);
                    resolve();
                }
            }
        });
        xmlhttp.open("POST", `/filter/${fileName.replace(/\//g, '$')}`, true);
        xmlhttp.send(); 
        return response;
    }

    const getChanges = (maxitems) => {
        var xmlhttp = new XMLHttpRequest();
        var response = new Promise(resolve => {
            xmlhttp.onreadystatechange = function () {
                if (this.readyState == 4 && this.status == 200) {
                    resolve(JSON.parse(this.responseText));
                }
            }
        });
        xmlhttp.open("GET", `/sort/changes?take=${maxitems}`, true);
        xmlhttp.send();
        return response;
    }

    const getCommits = (maxitems) => {
        var xmlhttp = new XMLHttpRequest();
        var response = new Promise(resolve => {
            xmlhttp.onreadystatechange = function () {
                if (this.readyState == 4 && this.status == 200) {
                    resolve(JSON.parse(this.responseText));
                }
            }
        });
        xmlhttp.open("GET", `/sort/commits?take=${maxitems}`, true);
        xmlhttp.send();
        return response;
    }

    const updateDate = (startDate, endDate) => {
        var xmlhttp = new XMLHttpRequest();
        var response = new Promise((resolve, reject) => {
            xmlhttp.onreadystatechange = function () {
                if (this.readyState == 4) {
                    if(this.status == 200){
                        notify("DATE_UPDATED");
                        resolve()
                    }
                    else reject(this.responseText)
                }
            }
        });
        xmlhttp.open("PUT", `/log?before=${endDate}&after=${startDate}`, true);
        xmlhttp.send(); 
        return response;
    }

    return {
        subscribe,
        addToFilter,
        removeFromFilter,
        getFilter,
        getChanges,
        getCommits,
        updateDate
    }
}