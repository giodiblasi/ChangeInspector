const initUpdater = () => {
    const dateSelector = document.getElementById("date-selector");
    const endDate = dateSelector.querySelector('#end');
    const startDate = dateSelector.querySelector('#start');

    const update = (startDate, endDate) => {
        var xmlhttp = new XMLHttpRequest();
        var response = new Promise((resolve, reject) => {
            xmlhttp.onreadystatechange = function () {
                if (this.readyState == 4) {
                    this.status == 200
                    ? resolve()
                    : reject(this.responseText)
                }
            }
        });
        xmlhttp.open("PUT", `/log?before=${endDate}&after=${startDate}`, true);
        xmlhttp.send(); 
        return response;
    }
    document
        .getElementById("update-button")
        .addEventListener('click',  (e)=>{
            
            update(startDate.value, endDate.value)
            .then(()=>dateSelector.dispatchEvent(new Event('data_updated')))
            .catch(err=>console.log(err))
            
        });
}