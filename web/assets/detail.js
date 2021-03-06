const initDetail = (api) => {
    const details ={};
    let selectedFile = '';
    google.charts.load('current', { 'packages': ['corechart'] });
    const getData = (fileName) => {
        const detail = details[fileName];
        if(detail) return Promise.resolve(detail);

        var xmlhttp = new XMLHttpRequest();
        var response = new Promise(resolve => {
            xmlhttp.onreadystatechange = function () {
                if (this.readyState == 4 && this.status == 200) {
                    const response = JSON.parse(this.responseText);
                    details[fileName]=response;
                    resolve(response);
                }
            }
        });
        xmlhttp.open("GET", `/files/${fileName.replace(/\//g, '$')}/detail`, true);
        xmlhttp.send(); 
        return response;
    }

    draw = (detail) => {
        var data = google.visualization.arrayToDataTable([
            ['Change', 'Total'],
            ['Adds', detail.TotalAdds],
            ['Remotions', detail.TotalRemotions]
        ]);

        var options = {
            title: 'Adds/Remotions'
        };

        var chart = new google.visualization.PieChart(document.getElementById('detail_changes'));

        chart.draw(data, options);
    }

    const updateInfo=(fileName, detail)=>{
        document.getElementById('detail_file_name').innerHTML = fileName;
        document.getElementById('detail_total_commits').innerHTML = `Total commits: ${detail.Commits.length}`;
        document.getElementById('detail_total_changes').innerHTML = `Total Changes: ${detail.TotalChanges}`;
    }

    const addFileToFilter = () => api.addToFilter(selectedFile);

    const clear = () => {
        document.getElementById('card_detail').hidden = true;
        document.getElementById('card_detail_empty').hidden = false;
    }

    api.subscribe("DATE_UPDATED", clear)

    
    return {
        update: async (fileName) => {
            selectedFile = fileName;
            const detail = await getData(fileName);
            draw(detail);
            updateInfo(fileName, detail);
            document.getElementById('card_detail').hidden = false;
            document.getElementById('card_detail_empty').hidden = true;
            const button = document.getElementById('add_filter');
            button.removeEventListener('click', addFileToFilter);
            button.addEventListener('click', addFileToFilter);

        },
        clear
    }
}