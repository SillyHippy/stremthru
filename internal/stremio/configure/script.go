package configure

import "html/template"

func GetScriptStoreTokenDescription(storeNameId, storeTokenId string) template.JS {
	return template.JS(`function onStoreNameChangeUpdateStoreTokenDescription() {
  const descByName = {
    "*": "API Key",
    "": "StremThru Basic Auth Token (base64 encoded) from <a href='https://github.com/SillyHippy/stremthru?tab=readme-ov-file#configuration' target='_blank'><code>STREMTHRU_PROXY_AUTH</code></a>",
		alldebrid: "AllDebrid <a href='https://alldebrid.com/apikeys' target='_blank'>API Key</a>",
		debridlink: "DebridLink <a href='https://debrid-link.com/webapp/apikey' target='_blank'>API Key</a>",
		easydebrid: "EasyDebrid API Key",
		offcloud: "Offcloud <a href='https://offcloud.com/#/account' target='_blank'>credential</a> in <code>email:password</code> format, e.g. <code>john.doe@example.com:secret-password</code>",
		premiumize: "Premiumize <a href='https://www.premiumize.me/account' target='_blank'>API Key</a>",
		realdebrid: "RealDebrid <a href='https://real-debrid.com/apitoken' target='_blank'>API Token</a>",
		torbox: "TorBox <a href='https://torbox.app/settings' target='_blank'>API Key</a>",
  };
  const nameElem = document.querySelector("#` + storeNameId + `");
  if (!nameElem) {
    return;
  }
  const tokenDescElem = document.querySelector("#` + storeTokenId + ` + small > span.description");
  if (tokenDescElem) {
    tokenDescElem.innerHTML = descByName[nameElem.value] || descByName["*"] || "";
  }
}

onStoreNameChangeUpdateStoreTokenDescription();

document.querySelector("#` + storeNameId + `").addEventListener("change", () => {
	onStoreNameChangeUpdateStoreTokenDescription();
});
`)
}
