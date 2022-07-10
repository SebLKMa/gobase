/*
sqlite> SELECT datetime(created_on, 'unixepoch', 'localtime'), temperature, humidity as created_on FROM metrics ORDER BY rowid DESC LIMIT 10;
2021-01-28 16:42:38|29.0900001525879|76.1600036621094
2021-01-27 15:13:21|29.6900005340576|69.6600036621094
2021-01-26 15:48:30|29.8299999237061|63.1500015258789
2021-01-26 15:47:25|29.7900009155273|63.1500015258789
2021-01-26 15:46:47|29.7900009155273|63.0800018310547
2021-01-26 15:46:17|29.7999992370605|63.25
2021-01-26 15:45:38|29.8700008392334|62.9900016784668
2021-01-26 15:42:47|29.8700008392334|63.0999984741211
2021-01-26 15:38:59|29.7199993133545|63.5099983215332
2021-01-26 15:37:41|29.7000007629395|63.4500007629395

Thermal 25.5
IAQ 15
Visual 8.25
Acoustics 26.4
Remainder 24.85
*/

// remainder will have no label
var components = ["Thermal Quality", "Indoor Air Quality", "Visual Quality", "Acoustic Quality", ""];
var scores = [25.5,15,8.25,26.4]; // the 4 IEQ scores
var total = scores.reduce((a, b) => a + b, 0) // iterate through the array, adding the current element value to the sum of the previous element values
var remainder = 100 - total
scores.push(remainder) // append to end of array


// text inside doughnut
// register plugin

Chart.pluginService.register({
  beforeDraw: function(chart) {
    var width = chart.chart.width,
        height = chart.chart.height,
        ctx = chart.chart.ctx;

    ctx.restore();
    var fontSize = (height / 114).toFixed(2);
    ctx.font = fontSize + "em sans-serif";
    ctx.textBaseline = "middle";

    var text = chart.options.centertext, // "75%",
        textX = Math.round((width - ctx.measureText(text).width) / 2);
    //var textY = height / 2 + (chart.titleBlock.height - 15); // when title display true

    var legendHeight =  chart.legend.height;
    var textY = height / 2 + legendHeight/2;
    //textY = textY + legendHeight/2; // when title and legend display are true

    ctx.fillText(text, textX, textY);
    ctx.save();
  }
}); 

// donut chart with remainder color transparent to background
new Chart(document.getElementById("doughnut-chart"), {
//var ctx = document.getElementById('doughnut-chart').getContext('2d');
//var myChart = new Chart(ctx, {
  type: 'doughnut',
  data: {
    labels: components,
    datasets: [
      {
        label: "Indoor Environment Quality (percentage)",
        backgroundColor: ["#3e95cd", "#8e5ea2","#c45850","#e8c3b9","#FFFFFF"],
        data: scores
      }
    ]
  },
  options: {
    responsive: true,
    title: {
      display: false,
      responsive: true,
      text: 'Indoor Environment Quality (percentage) on 2021-01-26'
    },
    legend: {
      display: true,
      responsive: true
    },
    centertext: total + "%"
  }
});


