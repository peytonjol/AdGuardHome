package dnsfilter

var safeSearchDomains = map[string]string{
	"yandex.com":     "213.180.193.56",
	"yandex.ru":      "213.180.193.56",
	"yandex.ua":      "213.180.193.56",
	"yandex.by":      "213.180.193.56",
	"yandex.kz":      "213.180.193.56",
	"www.yandex.com": "213.180.193.56",
	"www.yandex.ru":  "213.180.193.56",
	"www.yandex.ua":  "213.180.193.56",
	"www.yandex.by":  "213.180.193.56",
	"www.yandex.kz":  "213.180.193.56",
	"yandex.lt":  "213.180.193.56",
	"xmlsearch.yandex.ua":  "213.180.193.56",
	"www.yandex.lt":  "213.180.193.56",
	"yandex.net":  "213.180.193.56",
	"images.yandex.ru":  "213.180.193.56",
	"m.yandex.fr":  "213.180.193.56",
	"m.yandex.tm":  "213.180.193.56",
	"m.yandex.com.ge":  "213.180.193.56",
	"xmlsearch.yandex.com.tr":  "213.180.193.56",
	"m.yandex.co.il":  "213.180.193.56",
	"m.yandex.az":  "213.180.193.56",
	"yandex.tm":  "213.180.193.56",
	"people.yandex.kz":  "213.180.193.56",
	"family.yandex.com.tr":  "213.180.193.56",
	"m.yandex.kz":  "213.180.193.56",
	"xmlsearch.yandex.com":  "213.180.193.56",
	"www.yandex.kg":  "213.180.193.56",
	"www.yandex.uz":  "213.180.193.56",
	"play.yandex.com.tr":  "213.180.193.56",
	"yandex.co.il":  "213.180.193.56",
	"gorsel.yandex.com.tr":  "213.180.193.56",
	"images.yandex.com":  "213.180.193.56",
	"yandex.com.ge":  "213.180.193.56",
	"aile.yandex.com.tr":  "213.180.193.56",
	"yandex.kg":  "213.180.193.56",
	"m.yandex.ua":  "213.180.193.56",
	"m.yandex.lv":  "213.180.193.56",
	"game.yandex.com.tr":  "213.180.193.56",
	"yandex.com.tr":  "213.180.193.56",
	"www.yandex.com.am":  "213.180.193.56",
	"video.yandex.ua":  "213.180.193.56",
	"video.yandex.ru":  "213.180.193.56",
	"yandex.kz":  "213.180.193.56",
	"video.yandex.com.tr":  "213.180.193.56",
	"yandex.lv":  "213.180.193.56",
	"www.yandex.com.ge":  "213.180.193.56",
	"yandex.fr":  "213.180.193.56",
	"m.yandex.ru":  "213.180.193.56",
	"www.yandex.ua":  "213.180.193.56",
	"www.yandex.kz":  "213.180.193.56",
	"yandex.md":  "213.180.193.56",
	"games.yandex.com.tr":  "213.180.193.56",
	"www.yandex.co.il":  "213.180.193.56",
	"m.yandex.com":  "213.180.193.56",
	"yandex.ua":  "213.180.193.56",
	"yandex.by":  "213.180.193.56",
	"images.yandex.ua":  "213.180.193.56",
	"www.yandex.fr":  "213.180.193.56",
	"m.yandex.ee":  "213.180.193.56",
	"www.yandex.md":  "213.180.193.56",
	"www.yandex.ee":  "213.180.193.56",
	"people.yandex.com":  "213.180.193.56",
	"xmlsearch.yandex.kz":  "213.180.193.56",
	"yandex.ee":  "213.180.193.56",
	"video.yandex.com":  "213.180.193.56",
	"m.yandex.md":  "213.180.193.56",
	"video.yandex.by":  "213.180.193.56",
	"oyun.yandex.com.tr":  "213.180.193.56",
	"xmlsearch.yandex.ru":  "213.180.193.56",
	"www.yandex.lv":  "213.180.193.56",
	"m.yandex.uz":  "213.180.193.56",
	"www.yandex.tj":  "213.180.193.56",
	"people.yandex.by":  "213.180.193.56",
	"yandex.com.am":  "213.180.193.56",
	"people.yandex.ru":  "213.180.193.56",
	"images.yandex.kz":  "213.180.193.56",
	"people.yandex.com":  "213.180.193.56",
	"m.yandex.tj":  "213.180.193.56",
	"m.yandex.kg":  "213.180.193.56",
	"yandex.uz":  "213.180.193.56",
	"m.yandex.lt":  "213.180.193.56",
	"m.yandex.com.tr":  "213.180.193.56",
	"www.yandex.tm":  "213.180.193.56",
	"yandex.az":  "213.180.193.56",
	"m.yandex.com.am":  "213.180.193.56",
	"images.yandex.com.tr":  "213.180.193.56",
	"yandex.tj":  "213.180.193.56",
	"www.yandex.com.tr":  "213.180.193.56",
	"xmlsearch.yandex.by":  "213.180.193.56",
	"people.yandex.ua":  "213.180.193.56",
	"www.yandex.by":  "213.180.193.56",
	"video.yandex.kz":  "213.180.193.56",
	"www.yandex.az":  "213.180.193.56",
	"images.yandex.by":  "213.180.193.56",

	"www.bing.com": "204.79.197.220",

	"duckduckgo.com":       "54.241.17.246",
	"www.duckduckgo.com":   "54.241.17.246",
	"start.duckduckgo.com": "54.241.17.246",

	"www.google.com":    "forcesafesearch.google.com",
	"www.google.ad":     "forcesafesearch.google.com",
	"www.google.ae":     "forcesafesearch.google.com",
	"www.google.com.af": "forcesafesearch.google.com",
	"www.google.com.ag": "forcesafesearch.google.com",
	"www.google.com.ai": "forcesafesearch.google.com",
	"www.google.al":     "forcesafesearch.google.com",
	"www.google.am":     "forcesafesearch.google.com",
	"www.google.co.ao":  "forcesafesearch.google.com",
	"www.google.com.ar": "forcesafesearch.google.com",
	"www.google.as":     "forcesafesearch.google.com",
	"www.google.at":     "forcesafesearch.google.com",
	"www.google.com.au": "forcesafesearch.google.com",
	"www.google.az":     "forcesafesearch.google.com",
	"www.google.ba":     "forcesafesearch.google.com",
	"www.google.com.bd": "forcesafesearch.google.com",
	"www.google.be":     "forcesafesearch.google.com",
	"www.google.bf":     "forcesafesearch.google.com",
	"www.google.bg":     "forcesafesearch.google.com",
	"www.google.com.bh": "forcesafesearch.google.com",
	"www.google.bi":     "forcesafesearch.google.com",
	"www.google.bj":     "forcesafesearch.google.com",
	"www.google.com.bn": "forcesafesearch.google.com",
	"www.google.com.bo": "forcesafesearch.google.com",
	"www.google.com.br": "forcesafesearch.google.com",
	"www.google.bs":     "forcesafesearch.google.com",
	"www.google.bt":     "forcesafesearch.google.com",
	"www.google.co.bw":  "forcesafesearch.google.com",
	"www.google.by":     "forcesafesearch.google.com",
	"www.google.com.bz": "forcesafesearch.google.com",
	"www.google.ca":     "forcesafesearch.google.com",
	"www.google.cd":     "forcesafesearch.google.com",
	"www.google.cf":     "forcesafesearch.google.com",
	"www.google.cg":     "forcesafesearch.google.com",
	"www.google.ch":     "forcesafesearch.google.com",
	"www.google.ci":     "forcesafesearch.google.com",
	"www.google.co.ck":  "forcesafesearch.google.com",
	"www.google.cl":     "forcesafesearch.google.com",
	"www.google.cm":     "forcesafesearch.google.com",
	"www.google.cn":     "forcesafesearch.google.com",
	"www.google.com.co": "forcesafesearch.google.com",
	"www.google.co.cr":  "forcesafesearch.google.com",
	"www.google.com.cu": "forcesafesearch.google.com",
	"www.google.cv":     "forcesafesearch.google.com",
	"www.google.com.cy": "forcesafesearch.google.com",
	"www.google.cz":     "forcesafesearch.google.com",
	"www.google.de":     "forcesafesearch.google.com",
	"www.google.dj":     "forcesafesearch.google.com",
	"www.google.dk":     "forcesafesearch.google.com",
	"www.google.dm":     "forcesafesearch.google.com",
	"www.google.com.do": "forcesafesearch.google.com",
	"www.google.dz":     "forcesafesearch.google.com",
	"www.google.com.ec": "forcesafesearch.google.com",
	"www.google.ee":     "forcesafesearch.google.com",
	"www.google.com.eg": "forcesafesearch.google.com",
	"www.google.es":     "forcesafesearch.google.com",
	"www.google.com.et": "forcesafesearch.google.com",
	"www.google.fi":     "forcesafesearch.google.com",
	"www.google.com.fj": "forcesafesearch.google.com",
	"www.google.fm":     "forcesafesearch.google.com",
	"www.google.fr":     "forcesafesearch.google.com",
	"www.google.ga":     "forcesafesearch.google.com",
	"www.google.ge":     "forcesafesearch.google.com",
	"www.google.gg":     "forcesafesearch.google.com",
	"www.google.com.gh": "forcesafesearch.google.com",
	"www.google.com.gi": "forcesafesearch.google.com",
	"www.google.gl":     "forcesafesearch.google.com",
	"www.google.gm":     "forcesafesearch.google.com",
	"www.google.gp":     "forcesafesearch.google.com",
	"www.google.gr":     "forcesafesearch.google.com",
	"www.google.com.gt": "forcesafesearch.google.com",
	"www.google.gy":     "forcesafesearch.google.com",
	"www.google.com.hk": "forcesafesearch.google.com",
	"www.google.hn":     "forcesafesearch.google.com",
	"www.google.hr":     "forcesafesearch.google.com",
	"www.google.ht":     "forcesafesearch.google.com",
	"www.google.hu":     "forcesafesearch.google.com",
	"www.google.co.id":  "forcesafesearch.google.com",
	"www.google.ie":     "forcesafesearch.google.com",
	"www.google.co.il":  "forcesafesearch.google.com",
	"www.google.im":     "forcesafesearch.google.com",
	"www.google.co.in":  "forcesafesearch.google.com",
	"www.google.iq":     "forcesafesearch.google.com",
	"www.google.is":     "forcesafesearch.google.com",
	"www.google.it":     "forcesafesearch.google.com",
	"www.google.je":     "forcesafesearch.google.com",
	"www.google.com.jm": "forcesafesearch.google.com",
	"www.google.jo":     "forcesafesearch.google.com",
	"www.google.co.jp":  "forcesafesearch.google.com",
	"www.google.co.ke":  "forcesafesearch.google.com",
	"www.google.com.kh": "forcesafesearch.google.com",
	"www.google.ki":     "forcesafesearch.google.com",
	"www.google.kg":     "forcesafesearch.google.com",
	"www.google.co.kr":  "forcesafesearch.google.com",
	"www.google.com.kw": "forcesafesearch.google.com",
	"www.google.kz":     "forcesafesearch.google.com",
	"www.google.la":     "forcesafesearch.google.com",
	"www.google.com.lb": "forcesafesearch.google.com",
	"www.google.li":     "forcesafesearch.google.com",
	"www.google.lk":     "forcesafesearch.google.com",
	"www.google.co.ls":  "forcesafesearch.google.com",
	"www.google.lt":     "forcesafesearch.google.com",
	"www.google.lu":     "forcesafesearch.google.com",
	"www.google.lv":     "forcesafesearch.google.com",
	"www.google.com.ly": "forcesafesearch.google.com",
	"www.google.co.ma":  "forcesafesearch.google.com",
	"www.google.md":     "forcesafesearch.google.com",
	"www.google.me":     "forcesafesearch.google.com",
	"www.google.mg":     "forcesafesearch.google.com",
	"www.google.mk":     "forcesafesearch.google.com",
	"www.google.ml":     "forcesafesearch.google.com",
	"www.google.com.mm": "forcesafesearch.google.com",
	"www.google.mn":     "forcesafesearch.google.com",
	"www.google.ms":     "forcesafesearch.google.com",
	"www.google.com.mt": "forcesafesearch.google.com",
	"www.google.mu":     "forcesafesearch.google.com",
	"www.google.mv":     "forcesafesearch.google.com",
	"www.google.mw":     "forcesafesearch.google.com",
	"www.google.com.mx": "forcesafesearch.google.com",
	"www.google.com.my": "forcesafesearch.google.com",
	"www.google.co.mz":  "forcesafesearch.google.com",
	"www.google.com.na": "forcesafesearch.google.com",
	"www.google.com.nf": "forcesafesearch.google.com",
	"www.google.com.ng": "forcesafesearch.google.com",
	"www.google.com.ni": "forcesafesearch.google.com",
	"www.google.ne":     "forcesafesearch.google.com",
	"www.google.nl":     "forcesafesearch.google.com",
	"www.google.no":     "forcesafesearch.google.com",
	"www.google.com.np": "forcesafesearch.google.com",
	"www.google.nr":     "forcesafesearch.google.com",
	"www.google.nu":     "forcesafesearch.google.com",
	"www.google.co.nz":  "forcesafesearch.google.com",
	"www.google.com.om": "forcesafesearch.google.com",
	"www.google.com.pa": "forcesafesearch.google.com",
	"www.google.com.pe": "forcesafesearch.google.com",
	"www.google.com.pg": "forcesafesearch.google.com",
	"www.google.com.ph": "forcesafesearch.google.com",
	"www.google.com.pk": "forcesafesearch.google.com",
	"www.google.pl":     "forcesafesearch.google.com",
	"www.google.pn":     "forcesafesearch.google.com",
	"www.google.com.pr": "forcesafesearch.google.com",
	"www.google.ps":     "forcesafesearch.google.com",
	"www.google.pt":     "forcesafesearch.google.com",
	"www.google.com.py": "forcesafesearch.google.com",
	"www.google.com.qa": "forcesafesearch.google.com",
	"www.google.ro":     "forcesafesearch.google.com",
	"www.google.ru":     "forcesafesearch.google.com",
	"www.google.rw":     "forcesafesearch.google.com",
	"www.google.com.sa": "forcesafesearch.google.com",
	"www.google.com.sb": "forcesafesearch.google.com",
	"www.google.sc":     "forcesafesearch.google.com",
	"www.google.se":     "forcesafesearch.google.com",
	"www.google.com.sg": "forcesafesearch.google.com",
	"www.google.sh":     "forcesafesearch.google.com",
	"www.google.si":     "forcesafesearch.google.com",
	"www.google.sk":     "forcesafesearch.google.com",
	"www.google.com.sl": "forcesafesearch.google.com",
	"www.google.sn":     "forcesafesearch.google.com",
	"www.google.so":     "forcesafesearch.google.com",
	"www.google.sm":     "forcesafesearch.google.com",
	"www.google.sr":     "forcesafesearch.google.com",
	"www.google.st":     "forcesafesearch.google.com",
	"www.google.com.sv": "forcesafesearch.google.com",
	"www.google.td":     "forcesafesearch.google.com",
	"www.google.tg":     "forcesafesearch.google.com",
	"www.google.co.th":  "forcesafesearch.google.com",
	"www.google.com.tj": "forcesafesearch.google.com",
	"www.google.tk":     "forcesafesearch.google.com",
	"www.google.tl":     "forcesafesearch.google.com",
	"www.google.tm":     "forcesafesearch.google.com",
	"www.google.tn":     "forcesafesearch.google.com",
	"www.google.to":     "forcesafesearch.google.com",
	"www.google.com.tr": "forcesafesearch.google.com",
	"www.google.tt":     "forcesafesearch.google.com",
	"www.google.com.tw": "forcesafesearch.google.com",
	"www.google.co.tz":  "forcesafesearch.google.com",
	"www.google.com.ua": "forcesafesearch.google.com",
	"www.google.co.ug":  "forcesafesearch.google.com",
	"www.google.co.uk":  "forcesafesearch.google.com",
	"www.google.com.uy": "forcesafesearch.google.com",
	"www.google.co.uz":  "forcesafesearch.google.com",
	"www.google.com.vc": "forcesafesearch.google.com",
	"www.google.co.ve":  "forcesafesearch.google.com",
	"www.google.vg":     "forcesafesearch.google.com",
	"www.google.co.vi":  "forcesafesearch.google.com",
	"www.google.com.vn": "forcesafesearch.google.com",
	"www.google.vu":     "forcesafesearch.google.com",
	"www.google.ws":     "forcesafesearch.google.com",
	"www.google.rs":     "forcesafesearch.google.com",

	"www.youtube.com":          "restrictmoderate.youtube.com",
	"m.youtube.com":            "restrictmoderate.youtube.com",
	"youtubei.googleapis.com":  "restrictmoderate.youtube.com",
	"youtube.googleapis.com":   "restrictmoderate.youtube.com",
	"www.youtube-nocookie.com": "restrictmoderate.youtube.com",

	"pixabay.com": "safesearch.pixabay.com",
}
