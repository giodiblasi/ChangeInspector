const showReport = async (api) => {
    google.charts.load("current", { packages: ["corechart"] });
    const MAX_ITEMS = 10;
    function draw(title, dataLabels, element) {
        const barStyle = 'color: #76A7FA; stroke-width: 6';
        const columns =  [...dataLabels,{ role: 'style' }];
        var gdata = google.visualization.arrayToDataTable([columns, ["",0, barStyle]]);

        var options = {
            chart: {
                title,
            },
            animation: {
                duration: 700,
                easing: 'out',
                startup: true
            },
            bars: 'horizontal',
            height: 500
        };

        var chart = new google.visualization.BarChart(element);
        google.visualization.events.addListener(chart, 'select', selectHandler);
        function selectHandler(e) {
            const selectedItem = chart.getSelection()[0];
            if (selectedItem) {
                const selectedFileName = gdata.getValue(selectedItem.row, 0);
                var event = new CustomEvent('file_selected', { detail: selectedFileName });
                element.dispatchEvent(event);
            }
            else {
                element.dispatchEvent(new Event('clear_selection'));
            }
        }

        return {
            chart,
            update: (data) => {
                data = data.map(d=>[...d,barStyle]);
                gdata = google.visualization.arrayToDataTable([columns, ...data]);
                chart.draw(gdata, options);
            }
        }
    }

    const drawChanges = async (chartInfo) => {
        const data = (await api.getChanges(MAX_ITEMS));
        chartInfo.update(data);
    }

    const drawCommits = async (chartInfo) => {
        const data = (await api.getCommits(MAX_ITEMS));
        chartInfo.update(data);
    }

    const charts = await new Promise(resolve => google.charts.setOnLoadCallback(() => {
        const filesChart = draw(
            "File Changes",
            ["File", "Changes"],
            document.getElementById('chart_div_changes')
        );
        const commitsChart = draw(
            "File Commits",
            ["File", "Commits"],
            document.getElementById('chart_div_commits')
        )
        return resolve({
            filesChart,
            commitsChart
        });
    }));

    const update = ()=>{
        drawChanges(charts.filesChart);
        drawCommits(charts.commitsChart);
    }

    api.subscribe("FILTER_ITEM_ADDED", update)
    api.subscribe("FILTER_ITEM_REMOVED", update)
    api.subscribe("DATE_UPDATED", update)

    return {
        update
    };
}