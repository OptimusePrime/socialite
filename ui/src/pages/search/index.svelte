<!--
<script>
    // 1. Instantiate the search
    import instantsearch from "instantsearch.js";
    import { clearRefinements, hits, pagination, refinementList, searchBox } from "instantsearch.js/es/widgets";
    import { instantMeiliSearch } from "@meilisearch/instant-meilisearch";

    const search = instantsearch({
        indexName: "users",
        searchClient: instantMeiliSearch("http://localhost:7700"),
    });

    search.addWidgets([
        searchBox({
            // container: document.,
        }),
        clearRefinements({
            container: "#clear-refinements",
        }),
        refinementList({
            container: "#brand-list",
            attribute: "brand",
        }),
        hits({
            container: "#hits",
            templates: {
                item: `
        <div>
          <img src="{{image}}" align="left" alt="{{name}}" />
          <div class="hit-name">
            {{#helpers.highlight}}{ "attribute": "name" }{{/helpers.highlight}}
          </div>
          <div class="hit-description">
            {{#helpers.highlight}}{ "attribute": "description" }{{/helpers.highlight}}
          </div>
          <div class="hit-price">\${{price}}</div>
        </div>
      `,
            },
        }),
        pagination({
            container: "#pagination",
        }),
    ]);

    search.start();
</script>

<div class="ais-InstantSearch">
    <h1>InstantSearch.js e-commerce demo</h1>

    <div class="left-panel">
        <div id="clear-refinements"></div>

        <h2>Brands</h2>
        <div id="brand-list"></div>
    </div>

    <div class="right-panel">
        <div id="searchbox" class="ais-SearchBox"></div>
        <div id="hits"></div>
        <div id="pagination"></div>
    </div>
</div>
-->
