package repository_test

import (
	"testing"

	"github.com/tushar9989/github-search/internal/repository"
)

func TestSearch(t *testing.T) {
	html := getMockHTML()
	result, _ := repository.GetResultFromHTML(html)
	t.Run("CheckCount", func(t *testing.T) {
		expected := 6
		actual := len(result)
		if actual != expected {
			t.Error("Count does not match. Actual: ", actual)
		}
	})
	t.Run("CheckItemWithDescription", func(t *testing.T) {
		if "venky7974" != result[0].Owner {
			t.Error("Owner does not match. ", result[0].Owner)
		}
		if "demo_venky" != result[0].ProjectName {
			t.Error("ProjectName does not match. ", result[0].ProjectName)
		}
		if "zzzzzzzzzzzzzzzzzzzz" != result[0].Description {
			t.Error("Description does not match. ", result[0].Description)
		}
		if "2017-01-19T01:01:46Z" != result[0].LastUpdated {
			t.Error("LastUpdated does not match. ", result[0].LastUpdated)
		}
		if "https://github.com/venky7974/demo_venky" != result[0].RepositoryURL {
			t.Error("RepositoryURL does not match. ", result[0].RepositoryURL)
		}
	})
	t.Run("CheckItemWithNoDescription", func(t *testing.T) {
		if "" != result[2].Description {
			t.Error("Description is not empty. ", result[0].Description)
		}
	})
}

func getMockHTML() string {
	return `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
  <link rel="dns-prefetch" href="https://assets-cdn.github.com">
  <link rel="dns-prefetch" href="https://avatars0.githubusercontent.com">
  <link rel="dns-prefetch" href="https://avatars1.githubusercontent.com">
  <link rel="dns-prefetch" href="https://avatars2.githubusercontent.com">
  <link rel="dns-prefetch" href="https://avatars3.githubusercontent.com">
  <link rel="dns-prefetch" href="https://github-cloud.s3.amazonaws.com">
  <link rel="dns-prefetch" href="https://user-images.githubusercontent.com/">



  <link crossorigin="anonymous" media="all" integrity="sha512-hqbuBb0QOOmiWgl8a1V1N5q6TI/G0A2hVt/lCFYafR+fYsuXeRUcsdcb/yUyVEHYXktmUXl0Mx9s/BOUNZVq4w==" rel="stylesheet" href="https://assets-cdn.github.com/assets/frameworks-23c9e7262eee71bc6f67f6950190a162.css" />
  <link crossorigin="anonymous" media="all" integrity="sha512-eIYYJSKqUtZ1wca1cjvRpEFVDVFlRed6ZF1jtZ0E+plPCBYpZd78jUtSlGbFFouH9uhlhGRztTLDLwmNU6rL8w==" rel="stylesheet" href="https://assets-cdn.github.com/assets/github-50c729cf7e55c7554c6d0ceae2a0e938.css" />
  
  
  <link crossorigin="anonymous" media="all" integrity="sha512-cCY9KKeDzfGd+snVDZwcOIrTjK+JAFW4FR1c4OGJo1Z2QZNbAGVD6JlcnHL19LBcdByWHVxwrAPwigT/v/DWyQ==" rel="stylesheet" href="https://assets-cdn.github.com/assets/site-7472e7b4603d4095447d49d428375ab8.css" />
  

  <meta name="viewport" content="width=device-width">
  
  <title>Search · zzzzzzzzzzzzzzzzzzzz · GitHub</title>
    <meta name="description" content="GitHub is where people build software. More than 27 million people use GitHub to discover, fork, and contribute to over 80 million projects.">
    <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="GitHub">
  <link rel="fluid-icon" href="https://github.com/fluidicon.png" title="GitHub">
  <meta property="fb:app_id" content="1401488693436528">

    <meta property="og:url" content="https://github.com">
    <meta property="og:site_name" content="GitHub">
    <meta property="og:title" content="Build software better, together">
    <meta property="og:description" content="GitHub is where people build software. More than 27 million people use GitHub to discover, fork, and contribute to over 80 million projects.">
    <meta property="og:image" content="https://assets-cdn.github.com/images/modules/open_graph/github-logo.png">
    <meta property="og:image:type" content="image/png">
    <meta property="og:image:width" content="1200">
    <meta property="og:image:height" content="1200">
    <meta property="og:image" content="https://assets-cdn.github.com/images/modules/open_graph/github-mark.png">
    <meta property="og:image:type" content="image/png">
    <meta property="og:image:width" content="1200">
    <meta property="og:image:height" content="620">
    <meta property="og:image" content="https://assets-cdn.github.com/images/modules/open_graph/github-octocat.png">
    <meta property="og:image:type" content="image/png">
    <meta property="og:image:width" content="1200">
    <meta property="og:image:height" content="620">


  <link rel="assets" href="https://assets-cdn.github.com/">
  
  <meta name="pjax-timeout" content="1000">
  
  <meta name="request-id" content="E76C:4FA5:3B6C75D:7400E0E:5B0B0376" data-pjax-transient>


  

  <meta name="selected-link" value="code_search" data-pjax-transient>

    <meta name="google-site-verification" content="KT5gs8h0wvaagLKAVWq8bbeNwnZZK1r1XQysX3xurLU">
  <meta name="google-site-verification" content="ZzhVyEFwb7w3e0-uOTltm8Jsck2F5StVihD0exw2fsA">
  <meta name="google-site-verification" content="GXs5KoUUkNCoaAZn7wPN-t01Pywp9M3sEjnt_3_ZWPc">
    <meta name="google-analytics" content="UA-3769691-2">

<meta name="octolytics-host" content="collector.githubapp.com" /><meta name="octolytics-app-id" content="github" /><meta name="octolytics-event-url" content="https://collector.githubapp.com/github-external/browser_event" /><meta name="octolytics-dimension-request_id" content="E76C:4FA5:3B6C75D:7400E0E:5B0B0376" /><meta name="octolytics-dimension-region_edge" content="iad" /><meta name="octolytics-dimension-region_render" content="iad" />





  <meta class="js-ga-set" name="dimension1" content="Logged Out">


  

      <meta name="hostname" content="github.com">
    <meta name="user-login" content="">

      <meta name="expected-hostname" content="github.com">
    <meta name="js-proxy-site-detection-payload" content="OTI5YWQyODgyNjMwNGJjNjhkOGE0ZTY0NjNiOGIzY2FiMTU3YmVlMzMyMjUyNzNmMzdhMmNlZGMxZDc4NWNhNXx7InJlbW90ZV9hZGRyZXNzIjoiMTAzLjY2Ljk3LjExMyIsInJlcXVlc3RfaWQiOiJFNzZDOjRGQTU6M0I2Qzc1RDo3NDAwRTBFOjVCMEIwMzc2IiwidGltZXN0YW1wIjoxNTI3NDQ4NDM5LCJob3N0IjoiZ2l0aHViLmNvbSJ9">

    <meta name="enabled-features" content="UNIVERSE_BANNER,FREE_TRIALS,MARKETPLACE_INSIGHTS,MARKETPLACE_INSIGHTS_CONVERSION_PERCENTAGES,JUMP_TO">

  <meta name="html-safe-nonce" content="ee77791f172bd102555518bf00af2371d666fea5">

  <meta http-equiv="x-pjax-version" content="cb731f9d27da2ae3c5aaf1474ffe5d0a">
  

      <meta name="analytics-param-rename" content="q:ga-q" data-pjax-transient>




  <meta name="browser-stats-url" content="https://api.github.com/_private/browser/stats">

  <meta name="browser-errors-url" content="https://api.github.com/_private/browser/errors">

  <link rel="mask-icon" href="https://assets-cdn.github.com/pinned-octocat.svg" color="#000000">
  <link rel="icon" type="image/x-icon" class="js-site-favicon" href="https://assets-cdn.github.com/favicon.ico">

<meta name="theme-color" content="#1e2327">



<link rel="manifest" href="/manifest.json" crossOrigin="use-credentials">

  </head>

  <body class="logged-out env-production">
    

  <div class="position-relative js-header-wrapper ">
    <a href="#start-of-content" tabindex="1" class="px-2 py-4 bg-blue text-white show-on-focus js-skip-to-content">Skip to content</a>
    <div id="js-pjax-loader-bar" class="pjax-loader-bar"><div class="progress"></div></div>

    
    
    



        <header class="Header header-logged-out  position-relative f4 py-3" role="banner">
  <div class="container-lg d-flex px-3">
    <div class="d-flex flex-justify-between flex-items-center">
      <a class="header-logo-invertocat my-0" href="https://github.com/" aria-label="Homepage" data-ga-click="(Logged out) Header, go to homepage, icon:logo-wordmark">
        <svg height="32" class="octicon octicon-mark-github" viewBox="0 0 16 16" version="1.1" width="32" aria-hidden="true"><path fill-rule="evenodd" d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0 0 16 8c0-4.42-3.58-8-8-8z"/></svg>
      </a>

    </div>

    <div class="HeaderMenu HeaderMenu--bright d-flex flex-justify-between flex-auto">
        <nav class="mt-0">
          <ul class="d-flex list-style-none">
              <li class="ml-2">
                <a class="js-selected-navigation-item HeaderNavlink px-0 py-2 m-0" data-ga-click="Header, click, Nav menu - item:features" data-selected-links="/features /features/project-management /features/code-review /features/project-management /features/integrations /features" href="/features">
                  Features
</a>              </li>
              <li class="ml-4">
                <a class="js-selected-navigation-item HeaderNavlink px-0 py-2 m-0" data-ga-click="Header, click, Nav menu - item:business" data-selected-links="/business /business/security /business/customers /business" href="/business">
                  Business
</a>              </li>

              <li class="ml-4">
                <a class="js-selected-navigation-item HeaderNavlink px-0 py-2 m-0" data-ga-click="Header, click, Nav menu - item:explore" data-selected-links="/explore /trending /trending/developers /integrations /integrations/feature/code /integrations/feature/collaborate /integrations/feature/ship showcases showcases_search showcases_landing /explore" href="/explore">
                  Explore
</a>              </li>

              <li class="ml-4">
                    <a class="js-selected-navigation-item HeaderNavlink px-0 py-2 m-0" data-ga-click="Header, click, Nav menu - item:marketplace" data-selected-links=" /marketplace" href="/marketplace">
                      Marketplace
</a>              </li>
              <li class="ml-4">
                <a class="js-selected-navigation-item HeaderNavlink px-0 py-2 m-0" data-ga-click="Header, click, Nav menu - item:pricing" data-selected-links="/pricing /pricing/developer /pricing/team /pricing/business-hosted /pricing/business-enterprise /pricing" href="/pricing">
                  Pricing
</a>              </li>
          </ul>
        </nav>

      <div class="d-flex">
          <div class="d-lg-flex flex-items-center mr-3">
            <div class="header-search   js-site-search position-relative" role="search">
  <div class="position-relative">
    <!-- '" --><!-- </textarea></xmp> --></option></form><form class="js-site-search-form" data-unscoped-search-url="/search" action="/search" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" />
      <label class="form-control header-search-wrapper  js-chromeless-input-container">
        <input type="text"
          class="form-control header-search-input  js-site-search-focus "
          data-hotkey="s,/"
          name="q"
          value="zzzzzzzzzzzzzzzzzzzz"
          placeholder="Search GitHub"
          aria-label="Search GitHub"
          data-unscoped-placeholder="Search GitHub"
          data-scoped-placeholder="Search"
          autocapitalize="off"
          >
          <input type="hidden" class="js-site-search-type-field" name="type" >
      </label>
</form>  </div>
</div>

          </div>

        <span class="d-inline-block">
            <div class="HeaderNavlink px-0 py-2 m-0">
              <a class="text-bold text-white no-underline" href="/login?return_to=%2Fsearch%3Fq%3Dzzzzzzzzzzzzzzzzzzzz" data-ga-click="(Logged out) Header, clicked Sign in, text:sign-in">Sign in</a>
                <span class="text-gray">or</span>
                <a class="text-bold text-white no-underline" href="/join?source=header" data-ga-click="(Logged out) Header, clicked Sign up, text:sign-up">Sign up</a>
            </div>
        </span>
      </div>
    </div>
  </div>
</header>

  </div>

  <div id="start-of-content" class="show-on-focus"></div>

    <div id="js-flash-container">
</div>



  <div role="main" class="application-main ">
      
      <div id="js-pjax-container" data-pjax-container>
        

    <div class="container mt-4">
    <div class="columns">
      <div class="column one-fourth">
        <nav class="menu border" data-pjax>
            <a class="menu-item selected" href="/search?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Repositories">Repositories<span class="Counter ml-1 mt-1">6</span></a>
            <a class="menu-item " href="/search?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Code">Code<include-fragment src="https://github.com/search/count?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Code"></include-fragment></a>
            <a class="menu-item " href="/search?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Commits">Commits<include-fragment src="https://github.com/search/count?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Commits"></include-fragment></a>
            <a class="menu-item " href="/search?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Issues">Issues<include-fragment src="https://github.com/search/count?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Issues"></include-fragment></a>
            
            <a class="menu-item " href="/search?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Topics">Topics<include-fragment src="https://github.com/search/count?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Topics"></include-fragment></a>
            <a class="menu-item " href="/search?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Wikis">Wikis<include-fragment src="https://github.com/search/count?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Wikis"></include-fragment></a>
            <a class="menu-item " href="/search?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Users">Users<include-fragment src="https://github.com/search/count?q=zzzzzzzzzzzzzzzzzzzz&amp;type=Users"></include-fragment></a>
        </nav>
          <div class="border rounded-1 p-3 mb-3">
    <h2 class="d-inline-block f5 mb-2">
      Languages
    </h2>

    <ul class="filter-list small" data-pjax>
        <li>
            <span class="bar" style="width: 50%;"></span>

            <a class="filter-item" href="/search?l=HTML&amp;q=zzzzzzzzzzzzzzzzzzzz&amp;type=Repositories">
              <span class="count">1</span>
              HTML
</a>        </li>
        <li>
            <span class="bar" style="width: 50%;"></span>

            <a class="filter-item" href="/search?l=JavaScript&amp;q=zzzzzzzzzzzzzzzzzzzz&amp;type=Repositories">
              <span class="count">1</span>
              JavaScript
</a>        </li>
    </ul>
  </div>


        <p class="f6 mt-3">
          <a href="/search/advanced?q=zzzzzzzzzzzzzzzzzzzz">Advanced search</a>
          <button type="button" data-facebox="#search_cheatsheet_pane" data-facebox-class="shortcuts" class="btn-link ml-3">Cheat sheet</button>
        </p>
      </div>
      <div class="column three-fourths codesearch-results">
        <div class="pl-2">
          

  <div class="d-flex flex-justify-between border-bottom pb-3">
    <h3>
    6 repository results



</h3>

  </div>

          
  <ul class="repo-list">
      
<div class="repo-list-item d-flex flex-justify-start py-4 public source">
  <div class="col-8 pr-3">
    <h3>
      <a class="v-align-middle" data-hydro-click="{&quot;event_type&quot;:&quot;search_result.click&quot;,&quot;payload&quot;:{&quot;page_number&quot;:1,&quot;query&quot;:&quot;zzzzzzzzzzzzzzzzzzzz&quot;,&quot;result_position&quot;:1,&quot;click_id&quot;:79374446,&quot;result&quot;:{&quot;id&quot;:79374446,&quot;global_relay_id&quot;:&quot;MDEwOlJlcG9zaXRvcnk3OTM3NDQ0Ng==&quot;,&quot;model_name&quot;:&quot;Repository&quot;,&quot;url&quot;:&quot;https://github.com/venky7974/demo_venky&quot;},&quot;client_id&quot;:null,&quot;originating_request_id&quot;:&quot;E5B2:4FA5:3B67E67:73F7EB5:5B0B0304&quot;,&quot;user_id&quot;:null}}" data-hydro-click-hmac="b5373334769a06b6750070eb99ddf87d7793a8df1d8a3ed92613b7fd702fee7f" href="/venky7974/demo_venky">venky7974/demo_venky</a>

    </h3>

      <p class="col-9 d-inline-block text-gray mb-2 pr-4">
        <em>zzzzzzzzzzzzzzzzzzzz</em>
      </p>


    <div class="d-flex">

        <p class="f6 text-gray mr-3 mb-0 mt-2">
          Updated <relative-time datetime="2017-01-19T01:01:46Z">Jan 19, 2017</relative-time>
        </p>

    </div>
  </div>

  <div class="d-table-cell col-2 text-gray pt-2">
  </div>

</div>

      
<div class="repo-list-item d-flex flex-justify-start py-4 public source">
  <div class="col-8 pr-3">
    <h3>
      <a class="v-align-middle" data-hydro-click="{&quot;event_type&quot;:&quot;search_result.click&quot;,&quot;payload&quot;:{&quot;page_number&quot;:1,&quot;query&quot;:&quot;zzzzzzzzzzzzzzzzzzzz&quot;,&quot;result_position&quot;:2,&quot;click_id&quot;:107759147,&quot;result&quot;:{&quot;id&quot;:107759147,&quot;global_relay_id&quot;:&quot;MDEwOlJlcG9zaXRvcnkxMDc3NTkxNDc=&quot;,&quot;model_name&quot;:&quot;Repository&quot;,&quot;url&quot;:&quot;https://github.com/zzzzzzzzzzzzzzzzzzzzzzzzzzz111111/zzzzzzzzzzzzzzzzzzzzzzzzzzz11111.github.io&quot;},&quot;client_id&quot;:null,&quot;originating_request_id&quot;:&quot;E5B2:4FA5:3B67E67:73F7EB5:5B0B0304&quot;,&quot;user_id&quot;:null}}" data-hydro-click-hmac="28c7f155735afb80b5309470baf3e19592c1f1ae42ddd124bb556d056c1dde70" href="/zzzzzzzzzzzzzzzzzzzzzzzzzzz111111/zzzzzzzzzzzzzzzzzzzzzzzzzzz11111.github.io">zzzzzzzzzzzzzzzzzzzzzzzzzzz111111/<em>zzzzzzzzzzzzzzzzzzzzzzzzzzz11111</em>.github.io</a>

    </h3>



    <div class="d-flex">

        <p class="f6 text-gray mr-3 mb-0 mt-2">
          Updated <relative-time datetime="2017-10-21T07:03:21Z">Oct 21, 2017</relative-time>
        </p>

    </div>
  </div>

  <div class="d-table-cell col-2 text-gray pt-2">
  </div>

</div>

      
<div class="repo-list-item d-flex flex-justify-start py-4 public source">
  <div class="col-8 pr-3">
    <h3>
      <a class="v-align-middle" data-hydro-click="{&quot;event_type&quot;:&quot;search_result.click&quot;,&quot;payload&quot;:{&quot;page_number&quot;:1,&quot;query&quot;:&quot;zzzzzzzzzzzzzzzzzzzz&quot;,&quot;result_position&quot;:3,&quot;click_id&quot;:61920268,&quot;result&quot;:{&quot;id&quot;:61920268,&quot;global_relay_id&quot;:&quot;MDEwOlJlcG9zaXRvcnk2MTkyMDI2OA==&quot;,&quot;model_name&quot;:&quot;Repository&quot;,&quot;url&quot;:&quot;https://github.com/SuperMannnnnmn/Zzzzzzzzzzzzzzzzzzzz&quot;},&quot;client_id&quot;:null,&quot;originating_request_id&quot;:&quot;E5B2:4FA5:3B67E67:73F7EB5:5B0B0304&quot;,&quot;user_id&quot;:null}}" data-hydro-click-hmac="115bf9fdf074f8783a0361fe4d3a409ebaf3fe7678b8d228aa57c039f38d1197" href="/SuperMannnnnmn/Zzzzzzzzzzzzzzzzzzzz">SuperMannnnnmn/<em>Zzzzzzzzzzzzzzzzzzzz</em></a>

    </h3>



    <div class="d-flex">

        <p class="f6 text-gray mr-3 mb-0 mt-2">
          Updated <relative-time datetime="2016-06-25T00:55:18Z">Jun 25, 2016</relative-time>
        </p>

    </div>
  </div>

  <div class="d-table-cell col-2 text-gray pt-2">
  </div>

</div>

      
<div class="repo-list-item d-flex flex-justify-start py-4 public source">
  <div class="col-8 pr-3">
    <h3>
      <a class="v-align-middle" data-hydro-click="{&quot;event_type&quot;:&quot;search_result.click&quot;,&quot;payload&quot;:{&quot;page_number&quot;:1,&quot;query&quot;:&quot;zzzzzzzzzzzzzzzzzzzz&quot;,&quot;result_position&quot;:4,&quot;click_id&quot;:57118070,&quot;result&quot;:{&quot;id&quot;:57118070,&quot;global_relay_id&quot;:&quot;MDEwOlJlcG9zaXRvcnk1NzExODA3MA==&quot;,&quot;model_name&quot;:&quot;Repository&quot;,&quot;url&quot;:&quot;https://github.com/Mercle-Liang/zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz&quot;},&quot;client_id&quot;:null,&quot;originating_request_id&quot;:&quot;E5B2:4FA5:3B67E67:73F7EB5:5B0B0304&quot;,&quot;user_id&quot;:null}}" data-hydro-click-hmac="156495c2bed19747a4e7b7a420e8dc0ca75f498a1a0ec0241c6cd6bd1e95f8e5" href="/Mercle-Liang/zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz">Mercle-Liang/<em>zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz</em></a>

    </h3>



    <div class="d-flex">

        <p class="f6 text-gray mr-3 mb-0 mt-2">
          Updated <relative-time datetime="2016-04-26T10:11:55Z">Apr 26, 2016</relative-time>
        </p>

    </div>
  </div>

  <div class="d-table-cell col-2 text-gray pt-2">
        <span class="repo-language-color ml-0" style="background-color:#f1e05a;"></span>
      JavaScript
  </div>

</div>

      
<div class="repo-list-item d-flex flex-justify-start py-4 public source">
  <div class="col-8 pr-3">
    <h3>
      <a class="v-align-middle" data-hydro-click="{&quot;event_type&quot;:&quot;search_result.click&quot;,&quot;payload&quot;:{&quot;page_number&quot;:1,&quot;query&quot;:&quot;zzzzzzzzzzzzzzzzzzzz&quot;,&quot;result_position&quot;:5,&quot;click_id&quot;:105165803,&quot;result&quot;:{&quot;id&quot;:105165803,&quot;global_relay_id&quot;:&quot;MDEwOlJlcG9zaXRvcnkxMDUxNjU4MDM=&quot;,&quot;model_name&quot;:&quot;Repository&quot;,&quot;url&quot;:&quot;https://github.com/huihaitao/zzzzzzzzzzzzzzzzzzzzzzz&quot;},&quot;client_id&quot;:null,&quot;originating_request_id&quot;:&quot;E5B2:4FA5:3B67E67:73F7EB5:5B0B0304&quot;,&quot;user_id&quot;:null}}" data-hydro-click-hmac="57d4705f5e83c76130522f42408c03599c6018d39e7c2357370f65b1d8b5bb7c" href="/huihaitao/zzzzzzzzzzzzzzzzzzzzzzz">huihaitao/<em>zzzzzzzzzzzzzzzzzzzzzzz</em></a>

    </h3>



    <div class="d-flex">

        <p class="f6 text-gray mr-3 mb-0 mt-2">
          Updated <relative-time datetime="2017-09-28T15:31:54Z">Sep 28, 2017</relative-time>
        </p>

    </div>
  </div>

  <div class="d-table-cell col-2 text-gray pt-2">
        <span class="repo-language-color ml-0" style="background-color:#e34c26;"></span>
      HTML
  </div>

</div>

      
<div class="repo-list-item d-flex flex-justify-start py-4 public source">
  <div class="col-8 pr-3">
    <h3>
      <a class="v-align-middle" data-hydro-click="{&quot;event_type&quot;:&quot;search_result.click&quot;,&quot;payload&quot;:{&quot;page_number&quot;:1,&quot;query&quot;:&quot;zzzzzzzzzzzzzzzzzzzz&quot;,&quot;result_position&quot;:6,&quot;click_id&quot;:123333177,&quot;result&quot;:{&quot;id&quot;:123333177,&quot;global_relay_id&quot;:&quot;MDEwOlJlcG9zaXRvcnkxMjMzMzMxNzc=&quot;,&quot;model_name&quot;:&quot;Repository&quot;,&quot;url&quot;:&quot;https://github.com/pringleLips/zzzzzZZZZZzzZZZZZZZZ&quot;},&quot;client_id&quot;:null,&quot;originating_request_id&quot;:&quot;E5B2:4FA5:3B67E67:73F7EB5:5B0B0304&quot;,&quot;user_id&quot;:null}}" data-hydro-click-hmac="a398981c558ad6a1aa8a7709d885d07625f4bd3a02f8215af1221edea100b207" href="/pringleLips/zzzzzZZZZZzzZZZZZZZZ">pringleLips/<em>zzzzzZZZZZzzZZZZZZZZ</em></a>

    </h3>



    <div class="d-flex">

        <p class="f6 text-gray mr-3 mb-0 mt-2">
          Updated <relative-time datetime="2018-02-28T19:40:35Z">Feb 28, 2018</relative-time>
        </p>

    </div>
  </div>

  <div class="d-table-cell col-2 text-gray pt-2">
  </div>

</div>

  </ul>

        </div>
      </div>
    </div>
    <div id="search_cheatsheet_pane" class="instapaper_ignore readability-extra" style="display:none">
  <h2 class="facebox-header" data-facebox-id="facebox-header">Search cheat sheet <button type="button" class="btn-link js-see-all-search-cheatsheet"><small>(see all)</small></button></h2>

  <div class="markdown-body" data-facebox-id="facebox-description">

    <p>GitHub’s search supports a variety of different operations. Here’s a quick cheat sheet for some of the common searches.</p>
    <p>For more information, visit our <a href="https://help.github.com/articles/about-searching-on-github/">search help section</a>.</p>

    <!-- Basics -->
    <h2>Basic search</h2>
    <table>
      <thead>
        <tr>
          <th>This search</th>
          <th>Finds repositories with…</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>cat stars:&gt;100</td>
          <td>Find cat repositories with greater than 100 stars.</td>
        </tr>
        <tr>
          <td>user:defunkt</td>
          <td>Get all repositories from the user defunkt.</td>
        </tr>
        <tr>
          <td>tom location:"San Francisco, CA"</td>
          <td>Find all tom users in "San Francisco, CA".</td>
        </tr>
        <tr>
          <td>join extension:coffee</td>
          <td>Find all instances of join in code with coffee extension.</td>
        </tr>
         <tr>
          <td>NOT cat</td>
          <td>Excludes all results containing cat</td>
        </tr>
      </tbody>
    </table>

    <div class="js-more-cheatsheet-info d-none">
      <hr>

      <!-- Repositories -->
      <h2>Repository search
        <a href="https://help.github.com/articles/searching-repositories"><svg aria-label="Help" class="octicon octicon-link-external" viewBox="0 0 12 16" version="1.1" width="12" height="16" role="img"><path fill-rule="evenodd" d="M11 10h1v3c0 .55-.45 1-1 1H1c-.55 0-1-.45-1-1V3c0-.55.45-1 1-1h3v1H1v10h10v-3zM6 2l2.25 2.25L5 7.5 6.5 9l3.25-3.25L12 8V2H6z"/></svg></a>
      </h2>
      <p>Repository search looks through the projects you have access to on GitHub. You can also filter the results:</p>
      <table>
        <thead>
          <tr>
            <th>This search</th>
            <th>Finds repositories with…</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>cat stars:&gt;100</td>
            <td>Find cat repositories with greater than 100 stars.</td>
          </tr>
          <tr>
            <td>user:defunkt</td>
            <td>Get all repositories from the user defunkt.</td>
          </tr>
          <tr>
            <td>pugs pushed:&gt;2013-01-28</td>
            <td>Pugs repositories pushed to since Jan 28, 2013.</td>
          </tr>
          <tr>
            <td>node.js forks:&lt;200</td>
            <td>Find all node.js repositories with less than 200 forks.</td>
          </tr>
          <tr>
            <td>jquery size:1024..4089</td>
            <td>Find jquery repositories between the sizes 1024 and 4089 kB.</td>
          </tr>
          <tr>
            <td>gitx fork:true</td>
            <td>Repository search includes forks of gitx.</td>
          </tr>
          <tr>
            <td>gitx fork:only</td>
            <td>Repository search returns only forks of gitx.</td>
          </tr>
        </tbody>
      </table>

      <hr>

      <!-- Code -->
      <h2>Code search
        <a href="https://help.github.com/articles/searching-code"><svg aria-label="Help" class="octicon octicon-link-external" viewBox="0 0 12 16" version="1.1" width="12" height="16" role="img"><path fill-rule="evenodd" d="M11 10h1v3c0 .55-.45 1-1 1H1c-.55 0-1-.45-1-1V3c0-.55.45-1 1-1h3v1H1v10h10v-3zM6 2l2.25 2.25L5 7.5 6.5 9l3.25-3.25L12 8V2H6z"/></svg></a>
      </h2>
      <p>Code search looks through the files hosted on GitHub. You can also filter the results:</p>
      <table>
        <thead>
          <tr>
            <th>This search</th>
            <th>Finds repositories with…</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>install repo:charles/privaterepo</td>
            <td>Find all instances of install in code from the repository charles/privaterepo.</td>
          </tr>
          <tr>
            <td>shogun user:heroku</td>
            <td>Find references to shogun from all public heroku repositories.</td>
          </tr>
          <tr>
            <td>join extension:coffee</td>
            <td>Find all instances of join in code with coffee extension.</td>
          </tr>
          <tr>
            <td>system size:&gt;1000</td>
            <td>Find all instances of system in code of file size greater than 1000kbs.</td>
          </tr>
          <tr>
            <td>examples path:/docs/</td>
            <td>Find all examples in the path /docs/.</td>
          </tr>
          <tr>
            <td>replace fork:true</td>
            <td>Search replace in the source code of forks.</td>
          </tr>
        </tbody>
      </table>

      <hr>

      <!-- Issues -->
      <h2>Issue search
        <a href="https://help.github.com/articles/searching-issues"><svg aria-label="Help" class="octicon octicon-link-external" viewBox="0 0 12 16" version="1.1" width="12" height="16" role="img"><path fill-rule="evenodd" d="M11 10h1v3c0 .55-.45 1-1 1H1c-.55 0-1-.45-1-1V3c0-.55.45-1 1-1h3v1H1v10h10v-3zM6 2l2.25 2.25L5 7.5 6.5 9l3.25-3.25L12 8V2H6z"/></svg></a>
      </h2>
      <p>Issue search looks through issues and pull requests on GitHub. You can also filter the results:</p>
      <table>
        <thead>
          <tr>
            <th>This search</th>
            <th>Finds issues…</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>encoding user:heroku</td>
            <td>Encoding issues across the Heroku organization.</td>
          </tr>
          <tr>
            <td>cat is:open</td>
            <td>Find cat issues that are open.</td>
          </tr>
          <tr>
            <td>strange comments:&gt;42</td>
            <td>Issues with more than 42 comments.</td>
          </tr>
          <tr>
            <td>hard label:bug</td>
            <td>Hard issues labeled as a bug.</td>
          </tr>
          <tr>
            <td>author:mojombo</td>
            <td>All issues authored by mojombo.</td>
          </tr>
          <tr>
            <td>mentions:tpope</td>
            <td>All issues mentioning tpope.</td>
          </tr>
          <tr>
            <td>assignee:rtomayko</td>
            <td>All issues assigned to rtomayko.</td>
          </tr>
          <tr>
            <td>exception created:&gt;2012-12-31</td>
            <td>Created since the beginning of 2013.</td>
          </tr>
          <tr>
            <td>exception updated:&lt;2013-01-01</td>
            <td>Last updated before 2013.</td>
          </tr>
        </tbody>
      </table>

      <hr>

      <!-- Users -->
      <h2>User search
        <a href="https://help.github.com/articles/searching-users"><svg aria-label="Help" class="octicon octicon-link-external" viewBox="0 0 12 16" version="1.1" width="12" height="16" role="img"><path fill-rule="evenodd" d="M11 10h1v3c0 .55-.45 1-1 1H1c-.55 0-1-.45-1-1V3c0-.55.45-1 1-1h3v1H1v10h10v-3zM6 2l2.25 2.25L5 7.5 6.5 9l3.25-3.25L12 8V2H6z"/></svg></a>
      </h2>
      <p>User search finds users with an account on GitHub. You can also filter the results:</p>
      <table>
        <thead>
          <tr>
            <th>This search</th>
            <th>Finds repositories with…</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>fullname:"Linus Torvalds"</td>
            <td>Find users with the full name "Linus Torvalds".</td>
          </tr>
          <tr>
            <td>tom location:"San Francisco, CA"</td>
            <td>Find all tom users in "San Francisco, CA".</td>
          </tr>
          <tr>
            <td>chris followers:100..200</td>
            <td>Find all chris users with followers between 100 and 200.</td>
          </tr>
          <tr>
            <td>ryan repos:&gt;10</td>
            <td>Find all ryan users with more than 10 repositories.</td>
          </tr>
        </tbody>
      </table>
    </div>

  </div>
</div>

  </div>


      </div>
      <div class="modal-backdrop js-touch-events"></div>
  </div>

      
<div class="footer container-lg px-3" role="contentinfo">
  <div class="position-relative d-flex flex-justify-between pt-6 pb-2 mt-6 f6 text-gray border-top border-gray-light ">
    <ul class="list-style-none d-flex flex-wrap ">
      <li class="mr-3">&copy; 2018 <span title="0.19305s from unicorn-6df55b44bc-q7vkm">GitHub</span>, Inc.</li>
        <li class="mr-3"><a data-ga-click="Footer, go to terms, text:terms" href="https://github.com/site/terms">Terms</a></li>
        <li class="mr-3"><a data-ga-click="Footer, go to privacy, text:privacy" href="https://github.com/site/privacy">Privacy</a></li>
        <li class="mr-3"><a href="https://help.github.com/articles/github-security/" data-ga-click="Footer, go to security, text:security">Security</a></li>
        <li class="mr-3"><a href="https://status.github.com/" data-ga-click="Footer, go to status, text:status">Status</a></li>
        <li><a data-ga-click="Footer, go to help, text:help" href="https://help.github.com">Help</a></li>
    </ul>

    <a aria-label="Homepage" title="GitHub" class="footer-octicon" href="https://github.com">
      <svg height="24" class="octicon octicon-mark-github" viewBox="0 0 16 16" version="1.1" width="24" aria-hidden="true"><path fill-rule="evenodd" d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0 0 16 8c0-4.42-3.58-8-8-8z"/></svg>
</a>
   <ul class="list-style-none d-flex flex-wrap ">
        <li class="mr-3"><a data-ga-click="Footer, go to contact, text:contact" href="https://github.com/contact">Contact GitHub</a></li>
      <li class="mr-3"><a href="https://developer.github.com" data-ga-click="Footer, go to api, text:api">API</a></li>
      <li class="mr-3"><a href="https://training.github.com" data-ga-click="Footer, go to training, text:training">Training</a></li>
      <li class="mr-3"><a href="https://shop.github.com" data-ga-click="Footer, go to shop, text:shop">Shop</a></li>
        <li class="mr-3"><a href="https://blog.github.com" data-ga-click="Footer, go to blog, text:blog">Blog</a></li>
        <li><a data-ga-click="Footer, go to about, text:about" href="https://github.com/about">About</a></li>

    </ul>
  </div>
  <div class="d-flex flex-justify-center pb-6">
    <span class="f6 text-gray-light"></span>
  </div>
</div>



  <div id="ajax-error-message" class="ajax-error-message flash flash-error">
    <svg class="octicon octicon-alert" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"/></svg>
    <button type="button" class="flash-close js-ajax-error-dismiss" aria-label="Dismiss error">
      <svg class="octicon octicon-x" viewBox="0 0 12 16" version="1.1" width="12" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M7.48 8l3.75 3.75-1.48 1.48L6 9.48l-3.75 3.75-1.48-1.48L4.52 8 .77 4.25l1.48-1.48L6 6.52l3.75-3.75 1.48 1.48L7.48 8z"/></svg>
    </button>
    You can’t perform that action at this time.
  </div>


    <script crossorigin="anonymous" integrity="sha512-MWhu3P7Ss2c8wk/fvHCiNZ3LNVjHP4kaHYeO7xXrYQ1lt3MODm9jla1+FhjKVzoXk64jc05569Ia9K3Zq0Fd9Q==" type="application/javascript" src="https://assets-cdn.github.com/assets/compat-1c66c88316e8e9eacfac51af6ff01eac.js"></script>
    <script crossorigin="anonymous" integrity="sha512-X5Yz1+Qk9VsyD2X8vZkreaq36jjyMgMFXoVdiC7MDatN90E2HMy8SBo1Uv5+rHaLxCmF0icy+NwqIhLkci+Mvw==" type="application/javascript" src="https://assets-cdn.github.com/assets/frameworks-dca882fa8d6991b8dd62fde97105da60.js"></script>
    
    <script crossorigin="anonymous" async="async" integrity="sha512-kY82ZYHMFQQwb+2wSFyizzi88FEjErG6MGp1i9hD7RAZ1aMWNLy2v1whI+dzNVqINdxN2+AEeuuJH8WzWv4gag==" type="application/javascript" src="https://assets-cdn.github.com/assets/github-ee77c9e5b9dca5a57ed7eeeb5cf9abfd.js"></script>
    
    
    
  <div class="js-stale-session-flash stale-session-flash flash flash-warn flash-banner d-none">
    <svg class="octicon octicon-alert" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"/></svg>
    <span class="signed-in-tab-flash">You signed in with another tab or window. <a href="">Reload</a> to refresh your session.</span>
    <span class="signed-out-tab-flash">You signed out in another tab or window. <a href="">Reload</a> to refresh your session.</span>
  </div>
  <div class="facebox" id="facebox" style="display:none;">
  <div class="facebox-popup">
    <div class="facebox-content" role="dialog" aria-labelledby="facebox-header" aria-describedby="facebox-description">
    </div>
    <button type="button" class="facebox-close js-facebox-close" aria-label="Close modal">
      <svg class="octicon octicon-x" viewBox="0 0 12 16" version="1.1" width="12" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M7.48 8l3.75 3.75-1.48 1.48L6 9.48l-3.75 3.75-1.48-1.48L4.52 8 .77 4.25l1.48-1.48L6 6.52l3.75-3.75 1.48 1.48L7.48 8z"/></svg>
    </button>
  </div>
</div>

  <div class="Popover js-hovercard-content position-absolute" style="display: none; outline: none;" tabindex="0">
  <div class="Popover-message Popover-message--bottom-left Popover-message--large Box box-shadow-large" style="width:360px;">
  </div>
</div>

<div id="hovercard-aria-description" class="sr-only">
  Press h to open a hovercard with more details.
</div>


  </body>
</html>`
}
