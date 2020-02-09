let dateFormat = 'MMMM DD YYYY';
let data = []
let fetchData = new Request('/api/sensors');
let waterTimeInput


fetch(fetchData)
    .then(function (response) {
        if (!response.ok) {
            throw new Error('HTTP error, status = ' + response.status);
        }
        return response.json()
    }).then(function (payload) {
        payload.forEach(element => {
            data.push({
                t: moment(element.CreatedAt),
                y: (element.Temperature * 1.8) + 32
            })
        });
    });

function water(state) {
    var url = '/api/water/' + state;

    fetch(url, {
            method: 'POST'
        }).then(res => res.json())
        .then(response => console.log('Success:', JSON.stringify(response)))
        .catch(error => console.error('Error:', error));

}

window.onload = function () {
    waterTimeInput = document.getElementById("timeInput")
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
                        unit: 'minute'
                    }
                }],
                yAxes: [{
                    scaleLabel: {
                        display: true,
                        labelString: "Temperature (ºF)"
                    }
                }]
            }
        }
    });
}