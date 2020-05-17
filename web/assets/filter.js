const initFilter = () =>{
    const getData = () => {
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
                    resolve();
                }
            }
        });
        xmlhttp.open("DELETE", `/filter/${fileName.replace(/\//g, '$')}`, true);
        xmlhttp.send(); 
        return response;
    }


    const normalizeText = (text) => (text.length > 40) ? `...${text.substring(text.length-40, text.length)}` : text;

    return {
        update: async() => {
            const data = await getData();
            const list = document.getElementById("filter");
            data.map(file=>{
                const item = document.createElement("li");
                item.className = "list-group-item";
                
                const removeFromFilterButton = document.createElement("button");
                removeFromFilterButton.className = "btn btn-danger";
                removeFromFilterButton.innerHTML = "Remove";
                removeFromFilterButton.addEventListener('click', ()=>removeFromFilter(file));

                var name = document.createElement("p");
                name.innerHTML = `${normalizeText(file)}`;
                
                item.appendChild(name);
                item.appendChild(removeFromFilterButton);
                list.appendChild(item);
                
            },'');
        }
    }
}