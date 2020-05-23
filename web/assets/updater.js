const initUpdater = (api) => {
    const dateSelector = document.getElementById("date-selector");
    const endDate = dateSelector.querySelector('#end');
    const startDate = dateSelector.querySelector('#start');

    document
        .getElementById("update-button")
        .addEventListener('click',  (e)=>{
            api.updateDate(startDate.value, endDate.value)
        });
}