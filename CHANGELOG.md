# Changelog

## [0.1.0](https://github.com/SillyHippy/stremthru/compare/v0.28.1...0.1.0) (2025-01-05)


### Features

* add .env.example ([9f564f7](https://github.com/SillyHippy/stremthru/commit/9f564f760f2878cccb7e281f15dfa55adea1a667))
* add config for landing page ([e383a0d](https://github.com/SillyHippy/stremthru/commit/e383a0dc13233b5604571e980b443f1401931ea5))
* add Dockerfile ([ab4d4db](https://github.com/SillyHippy/stremthru/commit/ab4d4db4a0fbe1cbd302430e965370243752808e))
* add health/__debug__ endpoint ([94c4268](https://github.com/SillyHippy/stremthru/commit/94c4268b9d986b624a04647ff785a962e4d2da05))
* add landing page ([f499cbf](https://github.com/SillyHippy/stremthru/commit/f499cbf722da4c84057e87e326272b316245191d))
* add log for check manget ([cc8dafa](https://github.com/SillyHippy/stremthru/commit/cc8dafa20b23ed660f3e1d05b5f529ff30b14bfd))
* add support for uptream node ([f704542](https://github.com/SillyHippy/stremthru/commit/f70454298382413dfe7b04d92799eaf376173cd9))
* add version in health debug ([e9c1980](https://github.com/SillyHippy/stremthru/commit/e9c1980ac17786be8015b67c164ec26345a17dbd))
* allow api only store tunnel ([1867ed5](https://github.com/SillyHippy/stremthru/commit/1867ed5b83c8be070830b208c309b77855ed60cf))
* **buddy:** add auth token config ([f830911](https://github.com/SillyHippy/stremthru/commit/f8309119fb8a469662027961c413cb10a00bab1e))
* **buddy:** add local cache ([73b869e](https://github.com/SillyHippy/stremthru/commit/73b869ee810fd3547088794a4b500f55948ad755))
* **buddy:** forward client ip ([dee549b](https://github.com/SillyHippy/stremthru/commit/dee549b8970286ae24ccee56714227e7bc8ac60d))
* **buddy:** log packed error ([98a0ed1](https://github.com/SillyHippy/stremthru/commit/98a0ed1e69906a16dcb8dcefcb0867945e3edbae))
* **buddy:** send imdb id for movies ([48280b5](https://github.com/SillyHippy/stremthru/commit/48280b5af5e025c3dd08715d73fdb13467160577))
* **cache:** add method AddWithLifetime ([4bdf6b4](https://github.com/SillyHippy/stremthru/commit/4bdf6b43cc4d44fe0580956d732eb2dc0e406e9a))
* **core:** improve cache initialization ([8c31e5b](https://github.com/SillyHippy/stremthru/commit/8c31e5bc5123fa4ecd9e74c68c49423abbde50e6))
* **core:** rename magnet invalid error ([0b6be1f](https://github.com/SillyHippy/stremthru/commit/0b6be1f7b0264c878da5ce0a7464eda999f460f1))
* **db:** add 'heavy' tag for auto schema migration ([0d6b28f](https://github.com/SillyHippy/stremthru/commit/0d6b28fd4cf98cb27e0fc1940a560e06c1f31b59))
* **db:** add support for postgresql ([df3473c](https://github.com/SillyHippy/stremthru/commit/df3473c461f9aae8a95d56a715befbfbd6461a6f))
* **db:** handle connection and transaction better ([3c60920](https://github.com/SillyHippy/stremthru/commit/3c609203f1953ac545549d562729c9c74c3688d6))
* **db:** initial setup ([7371667](https://github.com/SillyHippy/stremthru/commit/73716677b9a9301763a61e5f584da13f489e65f9))
* **db:** log magnet_cache insert failure better ([0624b1a](https://github.com/SillyHippy/stremthru/commit/0624b1a8460f83e48cd5c94bba79d112438159ee))
* **db:** switch from libsql to sqlite3 ([dc9e0c8](https://github.com/SillyHippy/stremthru/commit/dc9e0c86212ea1c050348415c9f04c1feab10c1d))
* **db:** use wal mode for sqlite ([39f2b18](https://github.com/SillyHippy/stremthru/commit/39f2b18e8c628c01a00d9c933d1b9ed16f8cdc5f))
* decrease peer http timeout ([24cfdf4](https://github.com/SillyHippy/stremthru/commit/24cfdf40dc2ad6bea990c3c00e40c6d44c7afe96))
* extract request stuffs ([031dd77](https://github.com/SillyHippy/stremthru/commit/031dd77a09db0370e4344d682f9cd53e4c86a4d3))
* improve errors ([dcbe689](https://github.com/SillyHippy/stremthru/commit/dcbe689d057a0b1714d4fb68b245f3dd8d3a9fa7))
* initial implementation ([054a20f](https://github.com/SillyHippy/stremthru/commit/054a20f1ab84725f1221c9047c767db4d4db938a))
* log unexpected errors ([1d6e780](https://github.com/SillyHippy/stremthru/commit/1d6e7809250bbf6e60554e2c2cfb9b4e757aefaa))
* **magnet_cache:** adjust stale time ([e69f156](https://github.com/SillyHippy/stremthru/commit/e69f156491edd2d30e8b7ab3d406e2d036a44ba2))
* **magnet_cache:** extract and fix db stuffs ([8fbeafb](https://github.com/SillyHippy/stremthru/commit/8fbeafbb7c9a3099f203e1837d84adfbad9b0e2c))
* **magnet_cache:** get rid of unnecessary transaction ([fb7b244](https://github.com/SillyHippy/stremthru/commit/fb7b2441c1150d0499114183a891ffeabe1472c8))
* **magnet_cache:** preserve more specific sid for files ([0cc4e2d](https://github.com/SillyHippy/stremthru/commit/0cc4e2d488dfd4453d0d8748a8db08ff2e42f156))
* pass X-StremThru-Store-Name request header to response ([010f626](https://github.com/SillyHippy/stremthru/commit/010f62680aff744f41dddcb45c10325e5b7c41ac))
* **peer:** forward client ip ([13a555d](https://github.com/SillyHippy/stremthru/commit/13a555d5874e589856ce6e248ee665817b72b288))
* **peer:** introduce concept of peer ([18ced66](https://github.com/SillyHippy/stremthru/commit/18ced66a55e6a72630767e8231eeeb783011212d))
* print config at startup ([205f9e8](https://github.com/SillyHippy/stremthru/commit/205f9e8c7b32935aa4ddec7b9930e8c204ecb624))
* **request:** support passing query params ([f634ac7](https://github.com/SillyHippy/stremthru/commit/f634ac7861fdef1a71eb8185760c413737482d24))
* set http client timeout ([0803e1e](https://github.com/SillyHippy/stremthru/commit/0803e1e7e4a3349602b0b909d8577bfc2e8fb21b))
* store magnet cache info in db ([7b32556](https://github.com/SillyHippy/stremthru/commit/7b325560ac69f1a0e126df020feebccca8ca74c4))
* **store/easydebrid:** add initial implementation ([bd6f579](https://github.com/SillyHippy/stremthru/commit/bd6f5790f503d142192a28f5f43eccac0320065d))
* **store/offcloud:** add initial implementation ([a7156f5](https://github.com/SillyHippy/stremthru/commit/a7156f5d05195727814ae6a745a8929885504380))
* **store/premiumize:** improve magnet status detection ([81f1f2a](https://github.com/SillyHippy/stremthru/commit/81f1f2a07e220f605f1471d355b5d98a7ec41f14))
* **store/premiumize:** improve types ([6d92bd9](https://github.com/SillyHippy/stremthru/commit/6d92bd9c3d77381530f2ff02c6d63846a2586dfe))
* **store/realdebrid:** improve add magnet ([0e9a3ca](https://github.com/SillyHippy/stremthru/commit/0e9a3cab32d439e789fb99b54218530423570947))
* **store/realdebrid:** support passing client ip ([1265f1b](https://github.com/SillyHippy/stremthru/commit/1265f1bc8d1c897bfd27239d591f3793435eb751))
* **store/realdebrid:** update error codes ([8a49941](https://github.com/SillyHippy/stremthru/commit/8a499413edb220cfd32a124a2138797f1e7f28ad))
* **store:** add .hash for GetMagnet and ListMagnets ([aa93af5](https://github.com/SillyHippy/stremthru/commit/aa93af5f8fed38d2dd6ff8118e3d49893127cff6))
* **store:** add added_at field for magnet list/get ([d772158](https://github.com/SillyHippy/stremthru/commit/d772158053e4a519d43dc607125b4244afcc1ae0))
* **store:** add cache for torbox store ([dc2f26a](https://github.com/SillyHippy/stremthru/commit/dc2f26a8e8f3abec69504e8ea8a19b688688a0eb))
* **store:** add config for toggling tunnel ([b4ecc59](https://github.com/SillyHippy/stremthru/commit/b4ecc59cd95fbd990fed5fa14b66630bdd1fb5ad))
* **store:** add configurable user agent ([1e2d46c](https://github.com/SillyHippy/stremthru/commit/1e2d46cd9dcea23503139cfea051e592145c3be2))
* **store:** add enum for UserSubscriptionStatus ([5e9a0c9](https://github.com/SillyHippy/stremthru/commit/5e9a0c956b83695b05b0fb16f2b7bb58c483602b))
* **store:** add static videos ([ba0e514](https://github.com/SillyHippy/stremthru/commit/ba0e514f09aba83a02fd6bce7b628e1996a47214))
* **store:** add support for buddy ([5243279](https://github.com/SillyHippy/stremthru/commit/5243279eac80290843c2243223b5d3c9213afcb3))
* **store:** expose lowercase .hash for magnet ([709ec45](https://github.com/SillyHippy/stremthru/commit/709ec45a85892c2c8b98e2a9d8fce261f49ee1f6))
* **store:** improve magnet cache check ([f144494](https://github.com/SillyHippy/stremthru/commit/f144494b8d3e3f45d90d94eca81a18b5df88ce85))
* **store:** include filename in generated link ([85347ac](https://github.com/SillyHippy/stremthru/commit/85347acb1f31fe742f124110303c2391a89244d8))
* **store:** initial alldebrid integration ([8e80efe](https://github.com/SillyHippy/stremthru/commit/8e80efe5eb14b16277a4dce35b8de42ea3d965b6))
* **store:** initial debridlink integration ([c31f836](https://github.com/SillyHippy/stremthru/commit/c31f836f1bd7e0bba0ae6b62d6d64a7e90fb3a9d))
* **store:** initial premiumize integration ([d73fa42](https://github.com/SillyHippy/stremthru/commit/d73fa426f233d8d76875545c7162cb56bbd04f6c))
* **store:** initial realdebrid integration ([440cab2](https://github.com/SillyHippy/stremthru/commit/440cab237e5264df16b8655f2131f294a48cf5c0))
* **store:** initial torbox integration ([23d5cfd](https://github.com/SillyHippy/stremthru/commit/23d5cfdddb6dc06b58c56e2fff3ca4731914b60d))
* **store:** integrate buddy with all stores ([cd4998d](https://github.com/SillyHippy/stremthru/commit/cd4998d1543d72f17cdc14fca83082ea8216db0d))
* **store:** integrate upstream for check and track magnet cache ([f6bf4d7](https://github.com/SillyHippy/stremthru/commit/f6bf4d7900c58d61b238fb075c58725f5bd158bc))
* **store:** support json payload in request body ([aa73c7e](https://github.com/SillyHippy/stremthru/commit/aa73c7e58700cbeb26415c8b5de76ffd432ecd03))
* **store:** support pagination for list magnets ([0869539](https://github.com/SillyHippy/stremthru/commit/0869539a3ac4ac2e447af87658efc0612f05ec30))
* **store:** update error code for invalid store name ([281041e](https://github.com/SillyHippy/stremthru/commit/281041e978e2a6a66de37cc12b182df0b5ef9b4d))
* **store:** use local storage when buddy not available ([43fe0a3](https://github.com/SillyHippy/stremthru/commit/43fe0a308af600fd44df7a680ecafd5163c466f0))
* **stremio/sidekick:** add configure button for addons ([74c375a](https://github.com/SillyHippy/stremthru/commit/74c375adb10aed1f0678ea435a6c7bcb522a4339))
* **stremio/sidekick:** allow either id or name to change in reload ([18870ce](https://github.com/SillyHippy/stremthru/commit/18870cebceca4443364fd411d179076f16ded711))
* **stremio/sidekick:** improve button ux ([fc2928e](https://github.com/SillyHippy/stremthru/commit/fc2928efce1283182ba23216b9799b450c66e434))
* **stremio/sidekick:** improve login ux ([c4e3e34](https://github.com/SillyHippy/stremthru/commit/c4e3e34da9d45ef62f892656c176f1f3947730fd))
* **stremio/sidekick:** initial addon implementation ([1c88b4f](https://github.com/SillyHippy/stremthru/commit/1c88b4fbf3b6cfadbb8d12a0963eda807c404d21))
* **stremio/sidekick:** log ignored errors ([29bf96a](https://github.com/SillyHippy/stremthru/commit/29bf96adfbe82dc4ae87af2af586d3b9d93585bd))
* **stremio/sidekick:** show server error in ui ([580d42d](https://github.com/SillyHippy/stremthru/commit/580d42d455c9f65f63512ae1628884e5a6d6c209))
* **stremio/sidekick:** support disabling addon ([e0ea92d](https://github.com/SillyHippy/stremthru/commit/e0ea92d63cb74c4d475504d2d8ea3f662bd5b205))
* **stremio/sidekick:** support login with auth token ([a9f2967](https://github.com/SillyHippy/stremthru/commit/a9f29671c66bcc1a0c2fdabb6fd25fb5500ed423))
* **stremio/sidekick:** support reloading addon ([0f5c4a3](https://github.com/SillyHippy/stremthru/commit/0f5c4a343f65482c956c3cdcc771693ea08c2577))
* **stremio/sidekick:** update login ui ([1176f5a](https://github.com/SillyHippy/stremthru/commit/1176f5ade98fe6642ac3a8407485248fb973cec9))
* **stremio/store:** hide stremthru store if not usable ([c3a77ae](https://github.com/SillyHippy/stremthru/commit/c3a77ae6ae9e8666b6b21ec534fbda5d20bce243))
* **stremio/store:** log ignored errors ([9d3f641](https://github.com/SillyHippy/stremthru/commit/9d3f6419a40125851180364382837b00f1a93a89))
* **stremio/store:** update token field description ([2591e2f](https://github.com/SillyHippy/stremthru/commit/2591e2fbb2ee31197448cab2d43a43fc6b38ed47))
* **stremio/wrap:** add cache only config and sort ([472d937](https://github.com/SillyHippy/stremthru/commit/472d937cd5e1a730dfc09aa12d76c278b3b4228a))
* **stremio/wrap:** add configure button for upstream addon ([618e1df](https://github.com/SillyHippy/stremthru/commit/618e1df549906a8b40471e57a2f44f24c4ad5c3a))
* **stremio/wrap:** add easydebrid as store ([2f710b9](https://github.com/SillyHippy/stremthru/commit/2f710b9d29c81bb6c3a0779feb10011940d69640))
* **stremio/wrap:** add logs for static video redirects ([2231568](https://github.com/SillyHippy/stremthru/commit/2231568c1c24fdeb5d8974c6bad884187d9ebaa3))
* **stremio/wrap:** add store hint in addon name ([6b9040b](https://github.com/SillyHippy/stremthru/commit/6b9040b121a4e1ab5a8ddfbc14fb6a61a47a9327))
* **stremio/wrap:** auto-correct manifest url scheme ([5a23b0a](https://github.com/SillyHippy/stremthru/commit/5a23b0aa5c2a1b94f7955a522144099e3c3dffc2))
* **stremio/wrap:** forward client ip ([e8c4a5d](https://github.com/SillyHippy/stremthru/commit/e8c4a5dbc7956015cab80fff9f8c40f2589a42bb))
* **stremio/wrap:** forward client ip to upstream addon ([2ce883b](https://github.com/SillyHippy/stremthru/commit/2ce883bba0a13178405b707bfe5b3cac13a6f0e4))
* **stremio/wrap:** hide stremthru store if not usable ([b291197](https://github.com/SillyHippy/stremthru/commit/b29119718c05c7381375cc7294e5ca946ca2341f))
* **stremio/wrap:** improve magnet cache check ([b19fbf2](https://github.com/SillyHippy/stremthru/commit/b19fbf20d7495c30c7de4232c4811ac63bd8f463))
* **stremio/wrap:** log ignored errors ([1acbeac](https://github.com/SillyHippy/stremthru/commit/1acbeac3eafa547dc61f195b369a4eff171ec595))
* **stremio/wrap:** support magnet hash wrapping ([bd49120](https://github.com/SillyHippy/stremthru/commit/bd491209afecef5e50556961d42ac3e9d4fdb722))
* **stremio/wrap:** track added magnet ([e4a6c2b](https://github.com/SillyHippy/stremthru/commit/e4a6c2b8aa21418ba18a40a30a4e2af3ce076a6d))
* **stremio:** add addon - store ([39dca81](https://github.com/SillyHippy/stremthru/commit/39dca81be9eed485877fa8b3c85c9aca1930181c))
* **stremio:** add addon - wrap ([bfc9e1c](https://github.com/SillyHippy/stremthru/commit/bfc9e1c9ce1344568d6aae10edd097088780caeb))
* **stremio:** add config to toggle addons ([a6c06f8](https://github.com/SillyHippy/stremthru/commit/a6c06f811a0b6e45a3a5b7faae47503b3907e926))
* **stremio:** add description for addons ([e7f5758](https://github.com/SillyHippy/stremthru/commit/e7f5758972b0f50f415dde928d9b81104025a02b))
* **stremio:** add manifest validity check ([b619f10](https://github.com/SillyHippy/stremthru/commit/b619f10007b46f8ba018e9a6c4f8d4365f35f699))
* **stremio:** improve landing page for store ([7a0638c](https://github.com/SillyHippy/stremthru/commit/7a0638c48faed2f42757727764cbe578bd9a81ed))
* **stremio:** log packed error ([d16b617](https://github.com/SillyHippy/stremthru/commit/d16b6172d883d828120a24eec2c820241b59acde))
* **stremio:** mention store name in manifest ([559467c](https://github.com/SillyHippy/stremthru/commit/559467c76a249fb94abd9f94a3bbf8102fa2cddb))
* support fallback store auth config ([8a6cbd8](https://github.com/SillyHippy/stremthru/commit/8a6cbd8a89844d6b8f77b92f38e8668c1b644cce))
* support proxy auth ([9659c05](https://github.com/SillyHippy/stremthru/commit/9659c05c1629d2325664ff92500b1197c65ca426))
* support redis cache ([3bfbe70](https://github.com/SillyHippy/stremthru/commit/3bfbe70a7dfe16f12cb6689d5772e63eece4da8f))
* update header for buddy token ([45afd6d](https://github.com/SillyHippy/stremthru/commit/45afd6dd4cf5df37fa2f4c248acf1ff0ffd598f6))
* update log for track magnet ([f880962](https://github.com/SillyHippy/stremthru/commit/f88096275a8ddf49293a66711b7b59c4dd571be5))


### Bug Fixes

* **cache:** remove local cache for redis ([3bc1a20](https://github.com/SillyHippy/stremthru/commit/3bc1a200fcda36a3991c0c5aa1c5f144b325debe))
* **config:** handle env var with empty value ([7a1abcf](https://github.com/SillyHippy/stremthru/commit/7a1abcfee81d02fcf1c3128f4f2bb7733053f90e))
* core.ParseBasicAuth .Token value ([1cd1240](https://github.com/SillyHippy/stremthru/commit/1cd124069900a80e82e6d1969c8dfd8c3064179a))
* **core:** handle empty body for 204 status ([21417f1](https://github.com/SillyHippy/stremthru/commit/21417f11107ea7f3a412e2829d2aa0eef49eada6))
* **core:** remove empty dn query in ParseMagnet ([2aa59ff](https://github.com/SillyHippy/stremthru/commit/2aa59ffbc2e381d06487f35f78768ebb237e9080))
* **endpoint:** add missing early return ([274efee](https://github.com/SillyHippy/stremthru/commit/274efeea4fc338e016103a824c7c566e2d2d5bab))
* **endpoint:** do not send null for empty array ([93edc4d](https://github.com/SillyHippy/stremthru/commit/93edc4d99ea1d004f4d4aeb958385d53930f360c))
* **endpoint:** do not send null for empty array ([a2aba63](https://github.com/SillyHippy/stremthru/commit/a2aba633506284cb7e534fb4c057ea6536dcebc0))
* **endpoint:** expose delete magnet ([8171a29](https://github.com/SillyHippy/stremthru/commit/8171a29effe709a3d366a71c5896d98ecdafeb9d))
* **endpoint:** match landing page route exactly ([2d76b0e](https://github.com/SillyHippy/stremthru/commit/2d76b0ed7433cce78608cb62cd93dd4f969225d7))
* handle upstream service unavailable ([80d69ab](https://github.com/SillyHippy/stremthru/commit/80d69abc7266234c205eff726db300a0070e467d))
* **magnet_cache:** skip db query for empty input ([4a9cfdf](https://github.com/SillyHippy/stremthru/commit/4a9cfdf27e3764f25510685e3148e7c5cb1b239e))
* **peer_token:** fix schema file for postgresql ([aad8e7b](https://github.com/SillyHippy/stremthru/commit/aad8e7b37d504feaa749353d1937420b8607393b))
* **store/alldebrid:** ensure non-null .files for GetMagnet ([784ee1f](https://github.com/SillyHippy/stremthru/commit/784ee1fa5eb637b782c296a7c6e01e69224e815f))
* **store/debridlink:** handle not found for GetMagnet ([5cb1fb7](https://github.com/SillyHippy/stremthru/commit/5cb1fb7f20876e897a0794b904327dfe131fd831))
* **store/debridlink:** pass query params for ListSeedboxTorrents ([8a10e26](https://github.com/SillyHippy/stremthru/commit/8a10e26519108b899ad65072af2b01687a0a21d9))
* **store/premiumize:** handle not found for GetMagnet ([77dc312](https://github.com/SillyHippy/stremthru/commit/77dc31288f3bf084af2197dde5ed61506eca6e2b))
* **store/premiumize:** prefix file path with / ([a3eb584](https://github.com/SillyHippy/stremthru/commit/a3eb5844b78612d0f6beac25c8bf508924627545))
* **store/realdebrid:** deal with inconsistent type in response ([5f22bfb](https://github.com/SillyHippy/stremthru/commit/5f22bfb9d351619de0388a7c69bf58d4f3869b1a))
* **store/realdebrid:** fix type for .progress ([f5e5a0f](https://github.com/SillyHippy/stremthru/commit/f5e5a0f2f2f6c6b4f316866b6646f46c4dc3ce95))
* **store/torbox:** error handling for get magnet ([e28e401](https://github.com/SillyHippy/stremthru/commit/e28e401263ea0113dc5787ffad228eb640c0d82a))
* **store/torbox:** handle 404 for list torrents ([9730b8a](https://github.com/SillyHippy/stremthru/commit/9730b8a52bcba39a6c42180c2f2ce8f900b83441))
* **store/torbox:** handle extra item for list torrents ([a43167d](https://github.com/SillyHippy/stremthru/commit/a43167d10b12046e452b6a879d92318839e16b4b))
* **store:** nil-error for buddy ([1d597ab](https://github.com/SillyHippy/stremthru/commit/1d597ab7d4e2f440fe966e09b824c00c89bfd613))
* **store:** pass client ip only for non-proxy-authorized requests ([7f89bc3](https://github.com/SillyHippy/stremthru/commit/7f89bc3dd889100529e71b29b675f395c4fa7668))
* **store:** store name in error ([fee51a2](https://github.com/SillyHippy/stremthru/commit/fee51a26dcab67cd3cfd0ca5791906c2de3c3167))
* **stremio/sidekick:** fix double header on successful login ([62801c9](https://github.com/SillyHippy/stremthru/commit/62801c92ed161c60d5342de317ee09f63f56caff))
* **stremio/store:** allow trial subscription ([329781f](https://github.com/SillyHippy/stremthru/commit/329781fb493b106f413f301c1ec2f4759312c33c))
* **stremio/store:** check allowed method correctly ([3a97a23](https://github.com/SillyHippy/stremthru/commit/3a97a231aa93cc16af6005031acd9c7344efa3f1))
* **stremio/store:** handle empty metas ([2c91c21](https://github.com/SillyHippy/stremthru/commit/2c91c2164fc1e21cdd30a54968bdf334a0da0a45))
* **stremio/wrap:** handle missing .behaviorHints ([808136f](https://github.com/SillyHippy/stremthru/commit/808136fd0549925d3d7ff23774d8a8f8a1afc378))
* **stremio/wrap:** mark as wrapped only for proxy url ([d36a696](https://github.com/SillyHippy/stremthru/commit/d36a69658cfd8f2637189960440bd0f0f015731a))
* **stremio/wrap:** remove double link generation ([80606c5](https://github.com/SillyHippy/stremthru/commit/80606c50dd737d1621f93c14a8c3674a775cc798))
* **stremio:** add missing types for manifest ([0c25101](https://github.com/SillyHippy/stremthru/commit/0c251011a628ba6385855f4ff5d992ab7559df80))
* **stremio:** deduplicate id in configure page ([b56f934](https://github.com/SillyHippy/stremthru/commit/b56f9344a00478a1ec61820dd702b68aadee34aa))
* **stremio:** malformed manifest for store ([886d2a3](https://github.com/SillyHippy/stremthru/commit/886d2a342a40c0dd809cebe41d7688db73928f44))
* **stremio:** marshal json correctly for Resource ([8e0a196](https://github.com/SillyHippy/stremthru/commit/8e0a1964ebf4b680715e1ee977494a0a8c431cda))
* **stremio:** properly handle user data error for store ([ff6842d](https://github.com/SillyHippy/stremthru/commit/ff6842d74a51db2aa55ecffac2d79c51f95610cc))
* **stremio:** stop loading indicator on failed request ([e370417](https://github.com/SillyHippy/stremthru/commit/e3704178dfa013f91bb29ce02e62e793deafbcb3))
* **stremio:** use string body for addon client error ([87c764b](https://github.com/SillyHippy/stremthru/commit/87c764bf847e6b60591058681c9e00c78a0d5f96))
* update type to fix unexported field issue ([de5f5ba](https://github.com/SillyHippy/stremthru/commit/de5f5ba07be4aee6f24aa7917c6e7bca298d63b8))


### Performance Improvements

* **store:** cache access link token verification ([0db97d2](https://github.com/SillyHippy/stremthru/commit/0db97d2f8c235ce1f57ffa68e4db509bf645e0ef))


### Continuous Integration

* add release job ([d6bdd2e](https://github.com/SillyHippy/stremthru/commit/d6bdd2ea57153ae03483cb8bc6639ea04bd913cc))

## [0.28.1](https://github.com/SillyHippy/stremthru/compare/0.28.0...0.28.1) (2025-01-03)


### Bug Fixes

* **cache:** remove local cache for redis ([3bc1a20](https://github.com/SillyHippy/stremthru/commit/3bc1a200fcda36a3991c0c5aa1c5f144b325debe))

## [0.28.0](https://github.com/SillyHippy/stremthru/compare/0.27.1...0.28.0) (2025-01-03)


### Features

* **stremio/store:** hide stremthru store if not usable ([c3a77ae](https://github.com/SillyHippy/stremthru/commit/c3a77ae6ae9e8666b6b21ec534fbda5d20bce243))
* **stremio/wrap:** add cache only config and sort ([472d937](https://github.com/SillyHippy/stremthru/commit/472d937cd5e1a730dfc09aa12d76c278b3b4228a))
* **stremio/wrap:** add store hint in addon name ([6b9040b](https://github.com/SillyHippy/stremthru/commit/6b9040b121a4e1ab5a8ddfbc14fb6a61a47a9327))
* **stremio/wrap:** hide stremthru store if not usable ([b291197](https://github.com/SillyHippy/stremthru/commit/b29119718c05c7381375cc7294e5ca946ca2341f))


### Bug Fixes

* **magnet_cache:** skip db query for empty input ([4a9cfdf](https://github.com/SillyHippy/stremthru/commit/4a9cfdf27e3764f25510685e3148e7c5cb1b239e))
* **stremio/store:** handle empty metas ([2c91c21](https://github.com/SillyHippy/stremthru/commit/2c91c2164fc1e21cdd30a54968bdf334a0da0a45))

## [0.27.1](https://github.com/SillyHippy/stremthru/compare/0.27.0...0.27.1) (2025-01-03)


### Bug Fixes

* **stremio/wrap:** remove double link generation ([80606c5](https://github.com/SillyHippy/stremthru/commit/80606c50dd737d1621f93c14a8c3674a775cc798))

## [0.27.0](https://github.com/SillyHippy/stremthru/compare/0.26.0...0.27.0) (2025-01-03)


### Features

* **stremio/wrap:** add easydebrid as store ([2f710b9](https://github.com/SillyHippy/stremthru/commit/2f710b9d29c81bb6c3a0779feb10011940d69640))
* **stremio/wrap:** auto-correct manifest url scheme ([5a23b0a](https://github.com/SillyHippy/stremthru/commit/5a23b0aa5c2a1b94f7955a522144099e3c3dffc2))
* update log for track magnet ([f880962](https://github.com/SillyHippy/stremthru/commit/f88096275a8ddf49293a66711b7b59c4dd571be5))

## [0.26.0](https://github.com/SillyHippy/stremthru/compare/0.25.0...0.26.0) (2025-01-02)


### Features

* **store/easydebrid:** add initial implementation ([bd6f579](https://github.com/SillyHippy/stremthru/commit/bd6f5790f503d142192a28f5f43eccac0320065d))
* **stremio/sidekick:** allow either id or name to change in reload ([18870ce](https://github.com/SillyHippy/stremthru/commit/18870cebceca4443364fd411d179076f16ded711))


### Bug Fixes

* update type to fix unexported field issue ([de5f5ba](https://github.com/SillyHippy/stremthru/commit/de5f5ba07be4aee6f24aa7917c6e7bca298d63b8))

## [0.25.0](https://github.com/SillyHippy/stremthru/compare/0.24.0...0.25.0) (2025-01-02)


### Features

* allow api only store tunnel ([1867ed5](https://github.com/SillyHippy/stremthru/commit/1867ed5b83c8be070830b208c309b77855ed60cf))
* **buddy:** send imdb id for movies ([48280b5](https://github.com/SillyHippy/stremthru/commit/48280b5af5e025c3dd08715d73fdb13467160577))
* decrease peer http timeout ([24cfdf4](https://github.com/SillyHippy/stremthru/commit/24cfdf40dc2ad6bea990c3c00e40c6d44c7afe96))
* **magnet_cache:** preserve more specific sid for files ([0cc4e2d](https://github.com/SillyHippy/stremthru/commit/0cc4e2d488dfd4453d0d8748a8db08ff2e42f156))
* print config at startup ([205f9e8](https://github.com/SillyHippy/stremthru/commit/205f9e8c7b32935aa4ddec7b9930e8c204ecb624))

## [0.24.0](https://github.com/SillyHippy/stremthru/compare/0.23.0...0.24.0) (2025-01-02)


### Features

* **stremio:** log packed error ([d16b617](https://github.com/SillyHippy/stremthru/commit/d16b6172d883d828120a24eec2c820241b59acde))


### Bug Fixes

* **store/realdebrid:** fix type for .progress ([f5e5a0f](https://github.com/SillyHippy/stremthru/commit/f5e5a0f2f2f6c6b4f316866b6646f46c4dc3ce95))

## [0.23.0](https://github.com/SillyHippy/stremthru/compare/0.22.0...0.23.0) (2025-01-01)


### Features

* add log for check manget ([cc8dafa](https://github.com/SillyHippy/stremthru/commit/cc8dafa20b23ed660f3e1d05b5f529ff30b14bfd))
* **buddy:** log packed error ([98a0ed1](https://github.com/SillyHippy/stremthru/commit/98a0ed1e69906a16dcb8dcefcb0867945e3edbae))
* log unexpected errors ([1d6e780](https://github.com/SillyHippy/stremthru/commit/1d6e7809250bbf6e60554e2c2cfb9b4e757aefaa))
* **magnet_cache:** adjust stale time ([e69f156](https://github.com/SillyHippy/stremthru/commit/e69f156491edd2d30e8b7ab3d406e2d036a44ba2))
* set http client timeout ([0803e1e](https://github.com/SillyHippy/stremthru/commit/0803e1e7e4a3349602b0b909d8577bfc2e8fb21b))
* **store:** add configurable user agent ([1e2d46c](https://github.com/SillyHippy/stremthru/commit/1e2d46cd9dcea23503139cfea051e592145c3be2))
* **stremio/sidekick:** show server error in ui ([580d42d](https://github.com/SillyHippy/stremthru/commit/580d42d455c9f65f63512ae1628884e5a6d6c209))
* **stremio/sidekick:** support login with auth token ([a9f2967](https://github.com/SillyHippy/stremthru/commit/a9f29671c66bcc1a0c2fdabb6fd25fb5500ed423))
* **stremio/wrap:** track added magnet ([e4a6c2b](https://github.com/SillyHippy/stremthru/commit/e4a6c2b8aa21418ba18a40a30a4e2af3ce076a6d))

## [0.22.0](https://github.com/SillyHippy/stremthru/compare/0.21.0...0.22.0) (2024-12-31)


### Features

* **store:** add config for toggling tunnel ([b4ecc59](https://github.com/SillyHippy/stremthru/commit/b4ecc59cd95fbd990fed5fa14b66630bdd1fb5ad))
* **store:** improve magnet cache check ([f144494](https://github.com/SillyHippy/stremthru/commit/f144494b8d3e3f45d90d94eca81a18b5df88ce85))
* **stremio/wrap:** improve magnet cache check ([b19fbf2](https://github.com/SillyHippy/stremthru/commit/b19fbf20d7495c30c7de4232c4811ac63bd8f463))

## [0.21.0](https://github.com/SillyHippy/stremthru/compare/0.20.1...0.21.0) (2024-12-29)


### Features

* **stremio/sidekick:** log ignored errors ([29bf96a](https://github.com/SillyHippy/stremthru/commit/29bf96adfbe82dc4ae87af2af586d3b9d93585bd))
* **stremio/store:** log ignored errors ([9d3f641](https://github.com/SillyHippy/stremthru/commit/9d3f6419a40125851180364382837b00f1a93a89))
* **stremio/wrap:** add logs for static video redirects ([2231568](https://github.com/SillyHippy/stremthru/commit/2231568c1c24fdeb5d8974c6bad884187d9ebaa3))
* **stremio/wrap:** log ignored errors ([1acbeac](https://github.com/SillyHippy/stremthru/commit/1acbeac3eafa547dc61f195b369a4eff171ec595))

## [0.20.1](https://github.com/SillyHippy/stremthru/compare/0.20.0...0.20.1) (2024-12-28)


### Bug Fixes

* **stremio:** marshal json correctly for Resource ([8e0a196](https://github.com/SillyHippy/stremthru/commit/8e0a1964ebf4b680715e1ee977494a0a8c431cda))

## [0.20.0](https://github.com/SillyHippy/stremthru/compare/0.19.1...0.20.0) (2024-12-28)


### Features

* **buddy:** forward client ip ([dee549b](https://github.com/SillyHippy/stremthru/commit/dee549b8970286ae24ccee56714227e7bc8ac60d))
* **peer:** forward client ip ([13a555d](https://github.com/SillyHippy/stremthru/commit/13a555d5874e589856ce6e248ee665817b72b288))
* **stremio/wrap:** forward client ip ([e8c4a5d](https://github.com/SillyHippy/stremthru/commit/e8c4a5dbc7956015cab80fff9f8c40f2589a42bb))

## [0.19.1](https://github.com/SillyHippy/stremthru/compare/0.19.0...0.19.1) (2024-12-28)


### Bug Fixes

* **stremio/wrap:** handle missing .behaviorHints ([808136f](https://github.com/SillyHippy/stremthru/commit/808136fd0549925d3d7ff23774d8a8f8a1afc378))

## [0.19.0](https://github.com/SillyHippy/stremthru/compare/0.18.0...0.19.0) (2024-12-27)


### Features

* **stremio/sidekick:** add configure button for addons ([74c375a](https://github.com/SillyHippy/stremthru/commit/74c375adb10aed1f0678ea435a6c7bcb522a4339))
* **stremio/sidekick:** support reloading addon ([0f5c4a3](https://github.com/SillyHippy/stremthru/commit/0f5c4a343f65482c956c3cdcc771693ea08c2577))

## [0.18.0](https://github.com/SillyHippy/stremthru/compare/0.17.0...0.18.0) (2024-12-27)


### Features

* **stremio/wrap:** add configure button for upstream addon ([618e1df](https://github.com/SillyHippy/stremthru/commit/618e1df549906a8b40471e57a2f44f24c4ad5c3a))
* **stremio/wrap:** forward client ip to upstream addon ([2ce883b](https://github.com/SillyHippy/stremthru/commit/2ce883bba0a13178405b707bfe5b3cac13a6f0e4))
* **stremio:** add manifest validity check ([b619f10](https://github.com/SillyHippy/stremthru/commit/b619f10007b46f8ba018e9a6c4f8d4365f35f699))


### Bug Fixes

* **stremio/sidekick:** fix double header on successful login ([62801c9](https://github.com/SillyHippy/stremthru/commit/62801c92ed161c60d5342de317ee09f63f56caff))

## [0.17.0](https://github.com/SillyHippy/stremthru/compare/0.16.0...0.17.0) (2024-12-26)


### Features

* **cache:** add method AddWithLifetime ([4bdf6b4](https://github.com/SillyHippy/stremthru/commit/4bdf6b43cc4d44fe0580956d732eb2dc0e406e9a))
* **store:** add static videos ([ba0e514](https://github.com/SillyHippy/stremthru/commit/ba0e514f09aba83a02fd6bce7b628e1996a47214))
* **store:** include filename in generated link ([85347ac](https://github.com/SillyHippy/stremthru/commit/85347acb1f31fe742f124110303c2391a89244d8))
* **stremio/sidekick:** improve login ux ([c4e3e34](https://github.com/SillyHippy/stremthru/commit/c4e3e34da9d45ef62f892656c176f1f3947730fd))
* **stremio/store:** update token field description ([2591e2f](https://github.com/SillyHippy/stremthru/commit/2591e2fbb2ee31197448cab2d43a43fc6b38ed47))
* **stremio/wrap:** support magnet hash wrapping ([bd49120](https://github.com/SillyHippy/stremthru/commit/bd491209afecef5e50556961d42ac3e9d4fdb722))


### Bug Fixes

* core.ParseBasicAuth .Token value ([1cd1240](https://github.com/SillyHippy/stremthru/commit/1cd124069900a80e82e6d1969c8dfd8c3064179a))
* **stremio/store:** check allowed method correctly ([3a97a23](https://github.com/SillyHippy/stremthru/commit/3a97a231aa93cc16af6005031acd9c7344efa3f1))

## [0.16.0](https://github.com/SillyHippy/stremthru/compare/0.15.1...0.16.0) (2024-12-23)


### Features

* **request:** support passing query params ([f634ac7](https://github.com/SillyHippy/stremthru/commit/f634ac7861fdef1a71eb8185760c413737482d24))
* **store/offcloud:** add initial implementation ([a7156f5](https://github.com/SillyHippy/stremthru/commit/a7156f5d05195727814ae6a745a8929885504380))


### Bug Fixes

* **stremio/store:** allow trial subscription ([329781f](https://github.com/SillyHippy/stremthru/commit/329781fb493b106f413f301c1ec2f4759312c33c))

## [0.15.1](https://github.com/SillyHippy/stremthru/compare/0.15.0...0.15.1) (2024-12-21)


### Bug Fixes

* **stremio/wrap:** mark as wrapped only for proxy url ([d36a696](https://github.com/SillyHippy/stremthru/commit/d36a69658cfd8f2637189960440bd0f0f015731a))
* **stremio:** add missing types for manifest ([0c25101](https://github.com/SillyHippy/stremthru/commit/0c251011a628ba6385855f4ff5d992ab7559df80))

## [0.15.0](https://github.com/SillyHippy/stremthru/compare/0.14.0...0.15.0) (2024-12-21)


### Features

* **stremio/sidekick:** improve button ux ([fc2928e](https://github.com/SillyHippy/stremthru/commit/fc2928efce1283182ba23216b9799b450c66e434))
* **stremio/sidekick:** update login ui ([1176f5a](https://github.com/SillyHippy/stremthru/commit/1176f5ade98fe6642ac3a8407485248fb973cec9))
* **stremio:** add description for addons ([e7f5758](https://github.com/SillyHippy/stremthru/commit/e7f5758972b0f50f415dde928d9b81104025a02b))

## [0.14.0](https://github.com/SillyHippy/stremthru/compare/0.13.1...0.14.0) (2024-12-20)


### Features

* **stremio/sidekick:** initial addon implementation ([1c88b4f](https://github.com/SillyHippy/stremthru/commit/1c88b4fbf3b6cfadbb8d12a0963eda807c404d21))
* **stremio/sidekick:** support disabling addon ([e0ea92d](https://github.com/SillyHippy/stremthru/commit/e0ea92d63cb74c4d475504d2d8ea3f662bd5b205))


### Bug Fixes

* **stremio:** use string body for addon client error ([87c764b](https://github.com/SillyHippy/stremthru/commit/87c764bf847e6b60591058681c9e00c78a0d5f96))

## [0.13.1](https://github.com/SillyHippy/stremthru/compare/0.13.0...0.13.1) (2024-12-19)


### Bug Fixes

* **stremio:** deduplicate id in configure page ([b56f934](https://github.com/SillyHippy/stremthru/commit/b56f9344a00478a1ec61820dd702b68aadee34aa))
* **stremio:** stop loading indicator on failed request ([e370417](https://github.com/SillyHippy/stremthru/commit/e3704178dfa013f91bb29ce02e62e793deafbcb3))

## [0.13.0](https://github.com/SillyHippy/stremthru/compare/0.12.0...0.13.0) (2024-12-19)


### Features

* **stremio:** add addon - wrap ([bfc9e1c](https://github.com/SillyHippy/stremthru/commit/bfc9e1c9ce1344568d6aae10edd097088780caeb))

## [0.12.0](https://github.com/SillyHippy/stremthru/compare/0.11.1...0.12.0) (2024-12-19)


### Features

* add config for landing page ([e383a0d](https://github.com/SillyHippy/stremthru/commit/e383a0dc13233b5604571e980b443f1401931ea5))
* **stremio:** improve landing page for store ([7a0638c](https://github.com/SillyHippy/stremthru/commit/7a0638c48faed2f42757727764cbe578bd9a81ed))
* **stremio:** mention store name in manifest ([559467c](https://github.com/SillyHippy/stremthru/commit/559467c76a249fb94abd9f94a3bbf8102fa2cddb))


### Bug Fixes

* **endpoint:** match landing page route exactly ([2d76b0e](https://github.com/SillyHippy/stremthru/commit/2d76b0ed7433cce78608cb62cd93dd4f969225d7))

## [0.11.1](https://github.com/SillyHippy/stremthru/compare/0.11.0...0.11.1) (2024-12-18)


### Bug Fixes

* **stremio:** malformed manifest for store ([886d2a3](https://github.com/SillyHippy/stremthru/commit/886d2a342a40c0dd809cebe41d7688db73928f44))
* **stremio:** properly handle user data error for store ([ff6842d](https://github.com/SillyHippy/stremthru/commit/ff6842d74a51db2aa55ecffac2d79c51f95610cc))

## [0.11.0](https://github.com/SillyHippy/stremthru/compare/0.10.0...0.11.0) (2024-12-17)


### Features

* add landing page ([f499cbf](https://github.com/SillyHippy/stremthru/commit/f499cbf722da4c84057e87e326272b316245191d))
* add version in health debug ([e9c1980](https://github.com/SillyHippy/stremthru/commit/e9c1980ac17786be8015b67c164ec26345a17dbd))
* **stremio:** add addon - store ([39dca81](https://github.com/SillyHippy/stremthru/commit/39dca81be9eed485877fa8b3c85c9aca1930181c))
* **stremio:** add config to toggle addons ([a6c06f8](https://github.com/SillyHippy/stremthru/commit/a6c06f811a0b6e45a3a5b7faae47503b3907e926))

## [0.10.0](https://github.com/SillyHippy/stremthru/compare/0.9.0...0.10.0) (2024-12-15)


### Features

* **store:** add added_at field for magnet list/get ([d772158](https://github.com/SillyHippy/stremthru/commit/d772158053e4a519d43dc607125b4244afcc1ae0))

## [0.9.0](https://github.com/SillyHippy/stremthru/compare/0.8.0...0.9.0) (2024-12-10)


### Features

* **db:** add 'heavy' tag for auto schema migration ([0d6b28f](https://github.com/SillyHippy/stremthru/commit/0d6b28fd4cf98cb27e0fc1940a560e06c1f31b59))


### Bug Fixes

* **peer_token:** fix schema file for postgresql ([aad8e7b](https://github.com/SillyHippy/stremthru/commit/aad8e7b37d504feaa749353d1937420b8607393b))

## [0.8.0](https://github.com/SillyHippy/stremthru/compare/0.7.0...0.8.0) (2024-12-09)


### Features

* **db:** switch from libsql to sqlite3 ([dc9e0c8](https://github.com/SillyHippy/stremthru/commit/dc9e0c86212ea1c050348415c9f04c1feab10c1d))
* **magnet_cache:** get rid of unnecessary transaction ([fb7b244](https://github.com/SillyHippy/stremthru/commit/fb7b2441c1150d0499114183a891ffeabe1472c8))

## [0.7.0](https://github.com/SillyHippy/stremthru/compare/0.6.0...0.7.0) (2024-12-07)


### Features

* **db:** handle connection and transaction better ([3c60920](https://github.com/SillyHippy/stremthru/commit/3c609203f1953ac545549d562729c9c74c3688d6))
* **db:** log magnet_cache insert failure better ([0624b1a](https://github.com/SillyHippy/stremthru/commit/0624b1a8460f83e48cd5c94bba79d112438159ee))
* **magnet_cache:** extract and fix db stuffs ([8fbeafb](https://github.com/SillyHippy/stremthru/commit/8fbeafbb7c9a3099f203e1837d84adfbad9b0e2c))
* **store/realdebrid:** update error codes ([8a49941](https://github.com/SillyHippy/stremthru/commit/8a499413edb220cfd32a124a2138797f1e7f28ad))

## [0.6.0](https://github.com/SillyHippy/stremthru/compare/0.5.0...0.6.0) (2024-12-06)


### Features

* add support for uptream node ([f704542](https://github.com/SillyHippy/stremthru/commit/f70454298382413dfe7b04d92799eaf376173cd9))
* **db:** add support for postgresql ([df3473c](https://github.com/SillyHippy/stremthru/commit/df3473c461f9aae8a95d56a715befbfbd6461a6f))
* **db:** initial setup ([7371667](https://github.com/SillyHippy/stremthru/commit/73716677b9a9301763a61e5f584da13f489e65f9))
* **db:** use wal mode for sqlite ([39f2b18](https://github.com/SillyHippy/stremthru/commit/39f2b18e8c628c01a00d9c933d1b9ed16f8cdc5f))
* extract request stuffs ([031dd77](https://github.com/SillyHippy/stremthru/commit/031dd77a09db0370e4344d682f9cd53e4c86a4d3))
* **peer:** introduce concept of peer ([18ced66](https://github.com/SillyHippy/stremthru/commit/18ced66a55e6a72630767e8231eeeb783011212d))
* store magnet cache info in db ([7b32556](https://github.com/SillyHippy/stremthru/commit/7b325560ac69f1a0e126df020feebccca8ca74c4))
* **store:** integrate upstream for check and track magnet cache ([f6bf4d7](https://github.com/SillyHippy/stremthru/commit/f6bf4d7900c58d61b238fb075c58725f5bd158bc))
* **store:** update error code for invalid store name ([281041e](https://github.com/SillyHippy/stremthru/commit/281041e978e2a6a66de37cc12b182df0b5ef9b4d))
* update header for buddy token ([45afd6d](https://github.com/SillyHippy/stremthru/commit/45afd6dd4cf5df37fa2f4c248acf1ff0ffd598f6))


### Bug Fixes

* **config:** handle env var with empty value ([7a1abcf](https://github.com/SillyHippy/stremthru/commit/7a1abcfee81d02fcf1c3128f4f2bb7733053f90e))

## [0.5.0](https://github.com/SillyHippy/stremthru/compare/0.4.0...0.5.0) (2024-12-04)


### Features

* **store:** use local storage when buddy not available ([43fe0a3](https://github.com/SillyHippy/stremthru/commit/43fe0a308af600fd44df7a680ecafd5163c466f0))


### Bug Fixes

* **store:** pass client ip only for non-proxy-authorized requests ([7f89bc3](https://github.com/SillyHippy/stremthru/commit/7f89bc3dd889100529e71b29b675f395c4fa7668))

## [0.4.0](https://github.com/SillyHippy/stremthru/compare/0.3.0...0.4.0) (2024-12-03)


### Features

* **buddy:** add auth token config ([f830911](https://github.com/SillyHippy/stremthru/commit/f8309119fb8a469662027961c413cb10a00bab1e))
* **buddy:** add local cache ([73b869e](https://github.com/SillyHippy/stremthru/commit/73b869ee810fd3547088794a4b500f55948ad755))
* **core:** rename magnet invalid error ([0b6be1f](https://github.com/SillyHippy/stremthru/commit/0b6be1f7b0264c878da5ce0a7464eda999f460f1))
* **store/realdebrid:** support passing client ip ([1265f1b](https://github.com/SillyHippy/stremthru/commit/1265f1bc8d1c897bfd27239d591f3793435eb751))
* **store:** add support for buddy ([5243279](https://github.com/SillyHippy/stremthru/commit/5243279eac80290843c2243223b5d3c9213afcb3))
* **store:** integrate buddy with all stores ([cd4998d](https://github.com/SillyHippy/stremthru/commit/cd4998d1543d72f17cdc14fca83082ea8216db0d))
* support redis cache ([3bfbe70](https://github.com/SillyHippy/stremthru/commit/3bfbe70a7dfe16f12cb6689d5772e63eece4da8f))


### Bug Fixes

* handle upstream service unavailable ([80d69ab](https://github.com/SillyHippy/stremthru/commit/80d69abc7266234c205eff726db300a0070e467d))
* **store:** nil-error for buddy ([1d597ab](https://github.com/SillyHippy/stremthru/commit/1d597ab7d4e2f440fe966e09b824c00c89bfd613))

## [0.3.0](https://github.com/SillyHippy/stremthru/compare/0.2.0...0.3.0) (2024-11-24)


### Features

* improve errors ([dcbe689](https://github.com/SillyHippy/stremthru/commit/dcbe689d057a0b1714d4fb68b245f3dd8d3a9fa7))
* **store/premiumize:** improve types ([6d92bd9](https://github.com/SillyHippy/stremthru/commit/6d92bd9c3d77381530f2ff02c6d63846a2586dfe))

## [0.2.0](https://github.com/SillyHippy/stremthru/compare/0.1.0...0.2.0) (2024-11-23)


### Features

* **store:** support pagination for list magnets ([0869539](https://github.com/SillyHippy/stremthru/commit/0869539a3ac4ac2e447af87658efc0612f05ec30))


### Bug Fixes

* **store/torbox:** handle 404 for list torrents ([9730b8a](https://github.com/SillyHippy/stremthru/commit/9730b8a52bcba39a6c42180c2f2ce8f900b83441))
* **store/torbox:** handle extra item for list torrents ([a43167d](https://github.com/SillyHippy/stremthru/commit/a43167d10b12046e452b6a879d92318839e16b4b))

## 0.1.0 (2024-11-21)


### Features

* add .env.example ([9f564f7](https://github.com/SillyHippy/stremthru/commit/9f564f760f2878cccb7e281f15dfa55adea1a667))
* add Dockerfile ([ab4d4db](https://github.com/SillyHippy/stremthru/commit/ab4d4db4a0fbe1cbd302430e965370243752808e))
* add health/__debug__ endpoint ([94c4268](https://github.com/SillyHippy/stremthru/commit/94c4268b9d986b624a04647ff785a962e4d2da05))
* **core:** improve cache initialization ([8c31e5b](https://github.com/SillyHippy/stremthru/commit/8c31e5bc5123fa4ecd9e74c68c49423abbde50e6))
* initial implementation ([054a20f](https://github.com/SillyHippy/stremthru/commit/054a20f1ab84725f1221c9047c767db4d4db938a))
* pass X-StremThru-Store-Name request header to response ([010f626](https://github.com/SillyHippy/stremthru/commit/010f62680aff744f41dddcb45c10325e5b7c41ac))
* **store/premiumize:** improve magnet status detection ([81f1f2a](https://github.com/SillyHippy/stremthru/commit/81f1f2a07e220f605f1471d355b5d98a7ec41f14))
* **store/realdebrid:** improve add magnet ([0e9a3ca](https://github.com/SillyHippy/stremthru/commit/0e9a3cab32d439e789fb99b54218530423570947))
* **store:** add .hash for GetMagnet and ListMagnets ([aa93af5](https://github.com/SillyHippy/stremthru/commit/aa93af5f8fed38d2dd6ff8118e3d49893127cff6))
* **store:** add cache for torbox store ([dc2f26a](https://github.com/SillyHippy/stremthru/commit/dc2f26a8e8f3abec69504e8ea8a19b688688a0eb))
* **store:** add enum for UserSubscriptionStatus ([5e9a0c9](https://github.com/SillyHippy/stremthru/commit/5e9a0c956b83695b05b0fb16f2b7bb58c483602b))
* **store:** expose lowercase .hash for magnet ([709ec45](https://github.com/SillyHippy/stremthru/commit/709ec45a85892c2c8b98e2a9d8fce261f49ee1f6))
* **store:** initial alldebrid integration ([8e80efe](https://github.com/SillyHippy/stremthru/commit/8e80efe5eb14b16277a4dce35b8de42ea3d965b6))
* **store:** initial debridlink integration ([c31f836](https://github.com/SillyHippy/stremthru/commit/c31f836f1bd7e0bba0ae6b62d6d64a7e90fb3a9d))
* **store:** initial premiumize integration ([d73fa42](https://github.com/SillyHippy/stremthru/commit/d73fa426f233d8d76875545c7162cb56bbd04f6c))
* **store:** initial realdebrid integration ([440cab2](https://github.com/SillyHippy/stremthru/commit/440cab237e5264df16b8655f2131f294a48cf5c0))
* **store:** initial torbox integration ([23d5cfd](https://github.com/SillyHippy/stremthru/commit/23d5cfdddb6dc06b58c56e2fff3ca4731914b60d))
* **store:** support json payload in request body ([aa73c7e](https://github.com/SillyHippy/stremthru/commit/aa73c7e58700cbeb26415c8b5de76ffd432ecd03))
* support fallback store auth config ([8a6cbd8](https://github.com/SillyHippy/stremthru/commit/8a6cbd8a89844d6b8f77b92f38e8668c1b644cce))
* support proxy auth ([9659c05](https://github.com/SillyHippy/stremthru/commit/9659c05c1629d2325664ff92500b1197c65ca426))


### Bug Fixes

* **core:** handle empty body for 204 status ([21417f1](https://github.com/SillyHippy/stremthru/commit/21417f11107ea7f3a412e2829d2aa0eef49eada6))
* **core:** remove empty dn query in ParseMagnet ([2aa59ff](https://github.com/SillyHippy/stremthru/commit/2aa59ffbc2e381d06487f35f78768ebb237e9080))
* **endpoint:** add missing early return ([274efee](https://github.com/SillyHippy/stremthru/commit/274efeea4fc338e016103a824c7c566e2d2d5bab))
* **endpoint:** do not send null for empty array ([93edc4d](https://github.com/SillyHippy/stremthru/commit/93edc4d99ea1d004f4d4aeb958385d53930f360c))
* **endpoint:** do not send null for empty array ([a2aba63](https://github.com/SillyHippy/stremthru/commit/a2aba633506284cb7e534fb4c057ea6536dcebc0))
* **endpoint:** expose delete magnet ([8171a29](https://github.com/SillyHippy/stremthru/commit/8171a29effe709a3d366a71c5896d98ecdafeb9d))
* **store/alldebrid:** ensure non-null .files for GetMagnet ([784ee1f](https://github.com/SillyHippy/stremthru/commit/784ee1fa5eb637b782c296a7c6e01e69224e815f))
* **store/debridlink:** handle not found for GetMagnet ([5cb1fb7](https://github.com/SillyHippy/stremthru/commit/5cb1fb7f20876e897a0794b904327dfe131fd831))
* **store/debridlink:** pass query params for ListSeedboxTorrents ([8a10e26](https://github.com/SillyHippy/stremthru/commit/8a10e26519108b899ad65072af2b01687a0a21d9))
* **store/premiumize:** handle not found for GetMagnet ([77dc312](https://github.com/SillyHippy/stremthru/commit/77dc31288f3bf084af2197dde5ed61506eca6e2b))
* **store/premiumize:** prefix file path with / ([a3eb584](https://github.com/SillyHippy/stremthru/commit/a3eb5844b78612d0f6beac25c8bf508924627545))
* **store/realdebrid:** deal with inconsistent type in response ([5f22bfb](https://github.com/SillyHippy/stremthru/commit/5f22bfb9d351619de0388a7c69bf58d4f3869b1a))
* **store/torbox:** error handling for get magnet ([e28e401](https://github.com/SillyHippy/stremthru/commit/e28e401263ea0113dc5787ffad228eb640c0d82a))
* **store:** store name in error ([fee51a2](https://github.com/SillyHippy/stremthru/commit/fee51a26dcab67cd3cfd0ca5791906c2de3c3167))


### Performance Improvements

* **store:** cache access link token verification ([0db97d2](https://github.com/SillyHippy/stremthru/commit/0db97d2f8c235ce1f57ffa68e4db509bf645e0ef))


### Continuous Integration

* add release job ([d6bdd2e](https://github.com/SillyHippy/stremthru/commit/d6bdd2ea57153ae03483cb8bc6639ea04bd913cc))
