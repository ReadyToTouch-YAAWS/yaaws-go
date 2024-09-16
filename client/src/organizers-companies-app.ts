import {welcome} from "./welcome";

import urlStateContainer from "./framework/company_url_state_container";
import {
    COMPANY_SEARCH_QUERY,
    COMPANY_TYPE_CRITERIA_NAME,
    COMPANY_IN_FAVORITES_CRITERIA_NAME,
} from "./framework/company_criteria_names";
import {InputCheckboxes} from "./framework/checkboxes";
import {companyTypes} from "./framework/company_types";
import {htmlToNode} from "./framework/html";
import {Alias} from "./framework/alias";
import {renderSelected} from "./framework/selected_criteria";
import {firstQuerySelector} from "./framework/query_selector";
import {setStateByURLMapper} from "./framework/set_state_by_url";
import {toEnter} from "./framework/enter";

function markCompanyFavorite(companyId: number, favorite: boolean, callback: () => void) {
    fetch(`/api/v1/companies/${companyId}/favorite.json`, {
        method: "PATCH",
        body: JSON.stringify({
            favorite: favorite,
        }),
    }).then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = welcome();

            return;
        }

        callback();
    }).catch(console.error);
}

const $companies = document.querySelectorAll(".js-company");
const $resultCount = document.getElementById("js-result-count");

$companies.forEach(function ($company: HTMLElement) {
    const companyId = parseInt($company.getAttribute("data-company-id"));

    const $favorite = $company.querySelector(".js-company-favorite");
    $favorite.addEventListener("click", function () {
        const current = $favorite.classList.contains("in-favorite");
        const next = !current;

        markCompanyFavorite(companyId, next, function () {
            if (next) {
                $favorite.classList.add("in-favorite");

                $favorite.setAttribute("title", "Remove from favorites")
            } else {
                $favorite.classList.remove("in-favorite");

                $favorite.setAttribute("title", "Add to favorites")
            }
        });
    });
});


const $search = document.getElementById("js-company-query") as HTMLInputElement;
const $typeCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-company-type") as any as Array<HTMLInputElement>);
const $inFavoritesCheckbox = document.getElementById("js-criteria-in-favorites") as HTMLInputElement;
const $selectedCriteria = document.getElementById("js-company-selected-criteria");
const $reset = document.getElementById("js-criteria-reset");

$typeCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(COMPANY_TYPE_CRITERIA_NAME, state);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search();
});

$inFavoritesCheckbox.addEventListener("change", function () {
    urlStateContainer.setBoolCriteria(COMPANY_IN_FAVORITES_CRITERIA_NAME, $inFavoritesCheckbox.checked);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search();
});

const {
    setInputStateByURL,
    setCheckboxesStateByURL,
    setCheckboxStateByURL,
} = setStateByURLMapper(urlStateContainer);

function setStateByURL() {
    setInputStateByURL($search, COMPANY_SEARCH_QUERY);

    setCheckboxesStateByURL($typeCheckboxes, COMPANY_TYPE_CRITERIA_NAME);

    setCheckboxStateByURL($inFavoritesCheckbox, COMPANY_IN_FAVORITES_CRITERIA_NAME);
}

function renderSelectedCriteriaByURL() {
    const $views: Array<HTMLElement> = [];

    renderSelectedCheckboxes($views, COMPANY_TYPE_CRITERIA_NAME, companyTypes);
    renderSelectedCheckbox($views, COMPANY_IN_FAVORITES_CRITERIA_NAME, "Favorites");

    $selectedCriteria.innerHTML = "";
    $selectedCriteria.append(...$views);

    const visibility = $views.length === 0 ? "hidden" : "";
    $reset.style.visibility = visibility;
    $selectedCriteria.parentElement.style.visibility = visibility;
}

function renderSelectedCheckboxes(
    $views: Array<HTMLElement>,
    criteriaName: string,
    toAliases: (aliases: Array<string>) => Array<Alias>,
) {
    const aliases = urlStateContainer.getCriteria(criteriaName, []);

    toAliases(aliases).forEach(function (alias: Alias) {
        const $view = htmlToNode(renderSelected(alias.name));

        firstQuerySelector($view, "button").addEventListener("click", function () {
            urlStateContainer.removeAlias(criteriaName, alias.alias);
            urlStateContainer.setPage(1);
            urlStateContainer.storeCurrentState();

            updatePageState();
        });

        $views.push($view);
    });
}

function renderSelectedCheckbox($views: Array<HTMLElement>, criteriaName: string, title: string) {
    const checked = urlStateContainer.getCriteria(criteriaName, false);

    if (checked) {
        const $view = htmlToNode(renderSelected(title));

        firstQuerySelector($view, "button").addEventListener("click", function () {
            urlStateContainer.remove(criteriaName);
            urlStateContainer.setPage(1);
            urlStateContainer.storeCurrentState();

            updatePageState();
        });

        $views.push($view);
    }
}

const handleSearch = function () {
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    search();
}

$search.addEventListener("keyup", toEnter(handleSearch));
$search.addEventListener("change", function () {
    urlStateContainer.setStringCriteria(COMPANY_SEARCH_QUERY, $search.value);
});
$search.addEventListener("search", handleSearch);

$reset.addEventListener("click", function () {
    urlStateContainer.keep(COMPANY_SEARCH_QUERY);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    updatePageState();
});

function updatePageState() {
    setStateByURL();

    renderSelectedCriteriaByURL();

    search();
}

function search() {
    const query = $search.value.trim().toLowerCase();
    const types = urlStateContainer.getCriteria(COMPANY_TYPE_CRITERIA_NAME, []);
    const inFavorites = urlStateContainer.getCriteria(COMPANY_IN_FAVORITES_CRITERIA_NAME, false);

    const matchQuery = function ($company: HTMLElement): boolean {
        if (query.length === 0) {
            return true;
        }

        if ($company.getAttribute("data-company-name").toLowerCase().indexOf(query) !== -1) {
            return true;
        }

        if ($company.querySelector(".js-company-description").textContent.toLowerCase().indexOf(query) !== -1) {
            return true;
        }

        return false;
    }

    const match = function ($company: HTMLElement): boolean {
        if (!matchQuery($company)) {
            return false;
        }

        if (types.length > 0 && types.indexOf($company.getAttribute("data-company-type")) === -1) {
            return false;
        }

        if (inFavorites) {
            const $favorite = $company.querySelector(".js-company-favorite");
            const current = $favorite.classList.contains("in-favorite");

            if (!current) {
                return false;
            }
        }

        return true;
    }

    let total = 0;

    $companies.forEach(function ($company: HTMLElement) {
        if (match($company)) {
            $company.style.display = "block";

            total++;
        } else {
            $company.style.display = "none";
        }
    });

    $resultCount.innerHTML = total.toString();
}

updatePageState();

// <DEBUG>
/*
    function markCompanyFavoriteAll(favoriteGenerator: () => boolean) {
        document.querySelectorAll(".js-company").forEach(function ($element: HTMLElement) {
            markCompanyFavorite(
                parseInt($element.getAttribute("data-company-id")),
                favoriteGenerator(),
                function () {
                },
            );
        });
    }

    markCompanyFavoriteAll(function () {
        return Math.random() > 0.5;
    });
*/
// </DEBUG>