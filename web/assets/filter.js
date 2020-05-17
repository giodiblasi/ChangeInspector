const initFilter = (api) =>{
    const list = document.getElementById("filter");
    let itemnodes={}
    const addItem = (fileName) => {
        const item = document.createElement("li");
        item.className = "list-group-item";
        
        const removeFromFilterButton = document.createElement("button");
        removeFromFilterButton.className = "btn btn-danger";
        removeFromFilterButton.innerHTML = "Remove";
        removeFromFilterButton.addEventListener('click', ()=>api.removeFromFilter(fileName));

        var name = document.createElement("p");
        name.innerHTML = `${normalizeText(fileName)}`;
        
        item.appendChild(name);
        item.appendChild(removeFromFilterButton);
        list.appendChild(item);
        itemnodes[fileName] = item;
    }

    const deleteItem = (fileName) => {
        list.removeChild(itemnodes[fileName]);
    }

    const normalizeText = (text) => (text.length > 40) ? `...${text.substring(text.length-40, text.length)}` : text;
    
    api.subscribe("FILTER_ITEM_ADDED", addItem)
    api.subscribe("FILTER_ITEM_REMOVED", deleteItem)
    

    return {
        update: async() => {
            itemnodes={};
            const data = await api.getFilter();
            data.map(file=>addItem(file),'');
        }
    }
}