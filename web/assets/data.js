const showReport = (data) => {

    const dataArray = Object.keys(data).map(k => ({
        FileName: k,
        Info: data[k]
    }))

    google.charts.load('current', { 'packages': ['bar'] });

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
        chart.draw(gdata, google.charts.Bar.convertOptions(options));
    }

 
    //next: create sort data API 
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
        const data = (await getData('/sort/changes')).slice(10);
        google.charts.setOnLoadCallback(() => draw(
            data,
            "File Changes",
            ["File", "Changes"],
            document.getElementById('chart_div_changes')));
    }

    const drawCommits = async () => {
        const data  = (await getData('/sort/commits')).slice(10);
        google.charts.setOnLoadCallback(() => draw(
            data,
            "File Commits",
            ["File", "Commits"],
            document.getElementById('chart_div_commits')));
    }

    drawChanges();
    drawCommits();

}