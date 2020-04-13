const showReport = () => {

    google.charts.load('current', { 'packages': ['bar'] });
    const MAX_ITEMS = 10;
    function draw(dataArray, title, dataLabels, element) {
        var gdata = google.visualization.arrayToDataTable([
            dataLabels,
            ...dataArray
        ]);

        var options = {
            chart: {
                title,
            },
            bars: 'horizontal'
        };
        
        var chart = new google.charts.Bar(element);
        google.visualization.events.addListener(chart, 'select', selectHandler);
        function selectHandler(e) {
            const selectedItem = chart.getSelection()[0];
            if(selectedItem){
                const selectedFileName = gdata.getValue(selectedItem.row, 0);
                var event = new CustomEvent('file_selected', {detail: selectedFileName});
                element.dispatchEvent(event);
            }
            else{
                element.dispatchEvent(new Event('clear_selection'));
            }
        }

        chart.draw(gdata, google.charts.Bar.convertOptions(options));
    }

 
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
    const drawChanges = async () => {
        const data = (await getData(`/sort/changes?take=${MAX_ITEMS}`));
        google.charts.setOnLoadCallback(() => draw(
            data,
            "File Changes",
            ["File", "Changes"],
            document.getElementById('chart_div_changes')));
    }

    const drawCommits = async () => {
        const data  = (await getData(`/sort/commits?take=${MAX_ITEMS}`));
        google.charts.setOnLoadCallback(() => draw(
            data,
            "File Commits",
            ["File", "Commits"],
            document.getElementById('chart_div_commits')));
    }

    drawChanges();
    drawCommits();

}