let dateFormat = 'MMMM DD YYYY';
let data = [{
        t: moment('April 01 2017', dateFormat),
        y: 35
    },
    {
        t: moment('April 02 2017', dateFormat),
        y: 39
    },
    {
        t: moment('April 03 2017', dateFormat),
        y: 50
    },
    {
        t: moment('April 04 2017', dateFormat),
        y: 60
    },
    {
        t: moment('April 05 2017', dateFormat),
        y: 35
    },
    {
        t: moment('April 06 2017', dateFormat),
        y: 2
    },
    {
        t: moment('April 07 2017', dateFormat),
        y: 50
    },
    {
        t: moment('April 08 2017', dateFormat),
        y: 25
    }
]

window.onload = function () {
    var ctx = document.getElementById("stat-chart").getContext("2d");
    var myChart = new Chart(ctx, {
        type: "bar",
        data: {
            datasets: [{
                label: "Soil Temp",
                backgroundColor: 'rgb(255, 99, 132)',
                borderColor: 'rgb(255, 99, 132)',
                data: data,
                type: "line",
                pointRadius: 0,
                fill: false,
                lineTension: 0,
                borderWidth: 2
            }]
        },
        options: {
            scales: {
                xAxes: [{
                    type: "time",
                    distribution: "series",
                    ticks: {
                        source: "data",
                        autoSkip: true
                    },
                    time: {
                        unit: 'day'
                    }
                }],
                yAxes: [{
                    scaleLabel: {
                        display: true,
                        labelString: "Temperature (ºC)"
                    }
                }]
            }
        }
    });
}