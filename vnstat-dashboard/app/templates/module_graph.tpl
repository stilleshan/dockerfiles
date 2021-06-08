    <div class="container">
        <ul class="nav nav-tabs" id="graphTab" role="tablist">
            {if $jsonVersion gt 1}
                <li class="nav-item">
                    <a class="nav-link active" id="five-graph-tab" data-toggle="tab" href="#five-graph" role="tab" aria-controls="five-graph" aria-selected="true">Five Minute Graph</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" id="hourly-graph-tab" data-toggle="tab" href="#hourly-graph" role="tab" aria-controls="hourly-graph" aria-selected="false">Hourly Graph</a>
                </li>
            {else}
                <li class="nav-item">
                    <a class="nav-link active" id="hourly-graph-tab" data-toggle="tab" href="#hourly-graph" role="tab" aria-controls="hourly-graph" aria-selected="true">Hourly Graph</a>
                </li>
            {/if}
            <li class="nav-item">
                <a class="nav-link" id="daily-graph-tab" data-toggle="tab" href="#daily-graph" role="tab" aria-controls="daily-graph" aria-selected="false">Daily Graph</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" id="monthly-graph-tab" data-toggle="tab" href="#monthly-graph" role="tab" aria-controls="monthly-graph" aria-selected="false">Monthly Graph</a>
            </li>
        </ul>

        <div class="tab-content">
            {if $jsonVersion gt 1}
                <div class="tab-pane fade show active" id="five-graph" role="tabpanel" aria-labelledby="five-graph-tab">
                    <div id="fiveNetworkTrafficGraph" style="height: 300px;"></div>
                </div>

                <div class="tab-pane fade" id="hourly-graph" role="tabpanel" aria-labelledby="hourly-graph-tab">
                    <div id="hourlyNetworkTrafficGraph" style="height: 300px;"></div>
                </div>
            {else}
                <div class="tab-pane fade show active" id="hourly-graph" role="tabpanel" aria-labelledby="hourly-graph-tab">
                    <div id="hourlyNetworkTrafficGraph" style="height: 300px;"></div>
                </div>
            {/if}

            <div class="tab-pane fade" id="daily-graph" role="tabpanel" aria-labelledby="daily-graph-tab">
                <div id="dailyNetworkTrafficGraph" style="height: 300px;"></div>
            </div>

            <div class="tab-pane fade" id="monthly-graph" role="tabpanel" aria-labelledby="monthly-graph-tab">
                <div id="monthlyNetworkTrafficGraph" style="height: 300px;"></div>
            </div>
        </div>
    </div>
