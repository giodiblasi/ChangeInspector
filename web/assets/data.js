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

    const drawChanges = async () => {
        return new Promise(resolve => {
            const changesData = dataArray
                .sort((a, b) => b.Info.TotalChanges - a.Info.TotalChanges)
                .slice(10)
                .map(d => [d.FileName, d.Info.TotalChanges]);
            google.charts.setOnLoadCallback(() => draw(
                changesData,
                "File Changes",
                ["File", "Changes"],
                document.getElementById('chart_div_changes')));
            resolve();
        });
    }

    const drawCommits = async () => {
        return new Promise(resolve => {
            const commitsData = dataArray
                .sort((a, b) => b.Info.Commits.length - a.Info.Commits.length)
                .slice(10)
                .map(d => [d.FileName, d.Info.Commits.length]);


            google.charts.setOnLoadCallback(() => draw(
                commitsData,
                "File Commits",
                ["File", "Commits"],
                document.getElementById('chart_div_commits')));
            resolve();
        });
    }

    drawChanges();
    drawCommits();

}