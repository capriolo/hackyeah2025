<script>
    import { onDestroy, onMount } from "svelte";
    import Chart from "chart.js/auto";

    let { sleepData } = $props();

    const chartType = "doughnut";
    const showTitle = true;
    const showCenterText = true;
    const centerTextConfig = {
        mainText: null, // auto-generate if null
        subText: 'Całkowity czas snu',
        mainColor: '#fff',
        subColor: '#ddd',
        mainSize: 24,
        subSize: 14,
        mainWeight: 'bold',
        subWeight: 'normal'
    };
    let chart;

    const centerTextPlugin = {
        id: "centerText",
        beforeDraw(chart) {
            if (
                !chart.config.options.plugins.centerText ||
                !chart.config.options.plugins.centerText.display
            ) {
                return;
            }

            const ctx = chart.ctx;
            const centerX = (chart.chartArea.left + chart.chartArea.right) / 2;
            const centerY = (chart.chartArea.top + chart.chartArea.bottom) / 2;
            const options = chart.config.options.plugins.centerText;

            ctx.save();
            ctx.textAlign = "center";
            ctx.textBaseline = "middle";

            // Główny tekst
            if (options.mainText) {
                ctx.font = `${options.mainWeight} ${options.mainSize}px Arial`;
                ctx.fillStyle = options.mainColor;
                ctx.fillText(
                    options.mainText,
                    centerX,
                    centerY - options.mainSize / 3,
                );
            }

            // Dodatkowy tekst
            if (options.subText) {
                ctx.font = `${options.subWeight} ${options.subSize}px Arial`;
                ctx.fillStyle = options.subColor;
                ctx.fillText(
                    options.subText,
                    centerX,
                    centerY + options.mainSize / 3,
                );
            }

            ctx.restore();
        },
    };

    onMount(() => {
        if (!chartCanvas) return;

        const totalMinutes = sleepData.duration;
        const hours = Math.floor(totalMinutes / 60);
        const minutes = totalMinutes % 60;
        const autoMainText = `${hours}h ${minutes}m`;

        const config = {
            type: chartType,
            data: {
                labels: ["Sen Głęboki", "Sen Lekki", "Sen REM", "Wybudzony"],
                datasets: [
                    {
                        data: [
                            sleepData.deepSleepDuration,
                            sleepData.lightSleepDuration,
                            sleepData.remDuration,
                            sleepData.awakeDuration,
                        ],
                        backgroundColor: [
                            "#36A2EB",
                            "#4BC0C0",
                            "#FF6384",
                            "#FFCE56",
                        ],
                        borderWidth: 2,
                        borderColor: "#cb8766",
                    },
                ],
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    legend: {
                        position: "bottom",
                        labels: {
                            padding: 20,
                            color: "#fdfdfd",
                            font: {
                                size: 14,
                            },
                        },
                    },
                    centerText: {
                        display: showCenterText,
                        mainText: centerTextConfig.mainText || autoMainText,
                        subText: centerTextConfig.subText,
                        mainColor: centerTextConfig.mainColor,
                        subColor: centerTextConfig.subColor,
                        mainSize: centerTextConfig.mainSize,
                        subSize: centerTextConfig.subSize,
                        mainWeight: centerTextConfig.mainWeight,
                        subWeight: centerTextConfig.subWeight
                    },
                    tooltip: {
                        callbacks: {
                            label: function (context) {
                                const label = context.label || "";
                                const value = context.raw || 0;
                                const total = context.dataset.data.reduce(
                                    (a, b) => a + b,
                                    0,
                                );
                                const percentage = Math.round(
                                    (value / total) * 100,
                                );
                                return `${label}: ${value} min (${percentage}%)`;
                            },
                        },
                    },
                    title: {
                        display: showTitle,
                        text: `Analiza Snu - ${sleepData.date}`,
                        color: "#fff",
                        font: {
                            size: 16,
                        },
                    },
                },
            },
            plugins: [centerTextPlugin]
        };

        chart = new Chart(chartCanvas, config);

        chart.data.datasets[0].data = [
            sleepData.deepSleepDuration,
            sleepData.lightSleepDuration,
            sleepData.remDuration,
            sleepData.awakeDuration,
        ];

        if (showTitle) {
            chart.options.plugins.title.text = `Analiza Snu - ${sleepData.date}`;
        }

        chart.update();
    });

    onDestroy(() => {
        if (chart) {
            chart.destroy();
        }
    });


    // reload = (sleepData) => {
    //     console.log("reload", sleepData)
    //     chart.data.datasets[0].data = [
    //         sleepData.deepSleepDuration,
    //         sleepData.lightSleepDuration,
    //         sleepData.remDuration,
    //         sleepData.awakeDuration,
    //     ];
    //     chart.update();
    // }


    let chartCanvas = $state();
</script>

<div class="chart-container">
    <canvas bind:this={chartCanvas}></canvas>
</div>

<style>
    .chart-container {
        position: relative;
        width: 100%;
        height: 400px;
        max-width: 500px;
        margin: 0 auto;
    }
</style>
