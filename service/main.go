package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jorgeteixe/autodata/service/controllers"
)

func main() {
	r := mux.NewRouter()

	r.StrictSlash(true)

	gc := controllers.NewGeneralController()
	provc := controllers.NewProvinciasController()
	ac := controllers.NewAutoescuelasController()
	cec := controllers.NewCentrosExamenController()
	permc := controllers.NewPermisosController()
	tec := controllers.NewTiposExamenController()
	sc := controllers.NewSeccionesController()
	// rc := controllers.NewReportsController()

	r.HandleFunc("/ready", gc.ReadyHandler).Methods("GET")

	r.HandleFunc("/provincias", provc.GetProvinciasHandler).Methods("GET")
	r.HandleFunc("/provincias/{id}", provc.GetProvinciaHandler).Methods("GET")

	r.HandleFunc("/centros-examen", cec.GetCentrosExamenHandler).Methods("GET")
	r.HandleFunc("/centros-examen/{id}", cec.GetCentroExamenHandler).Methods("GET")

	r.HandleFunc("/autoescuelas", ac.GetAutoescuelasHandler).Methods("GET")
	r.HandleFunc("/autoescuelas/{id}", ac.GetAutoescuelaHandler).Methods("GET")

	r.HandleFunc("/secciones", sc.GetSeccionesHandler).Methods("GET")
	r.HandleFunc("/secciones/{id}", sc.GetSeccionHandler).Methods("GET")

	r.HandleFunc("/permisos", permc.GetPermisosHandler).Methods("GET")
	r.HandleFunc("/permisos/{id}", permc.GetPermisoHandler).Methods("GET")

	r.HandleFunc("/tipos-examen", tec.GetTiposExamenHandler).Methods("GET")
	r.HandleFunc("/tipos-examen/{id}", tec.GetTipoExamenHandler).Methods("GET")

	// r.HandleFunc("/reports", rc.GetReportsHandler).Methods("GET")
	// r.HandleFunc("/reports/{id}", rc.GetReportHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8090", logRequest(r)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", requestGetRemoteAddress(r), r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

// Request.RemoteAddress contains port, which we want to remove i.e.:
// "[::1]:58292" => "[::1]"
func ipAddrFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

// requestGetRemoteAddress returns ip address of the client making the request,
// taking into account http proxies
func requestGetRemoteAddress(r *http.Request) string {
	hdr := r.Header
	hdrRealIP := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")
	if hdrRealIP == "" && hdrForwardedFor == "" {
		return ipAddrFromRemoteAddr(r.RemoteAddr)
	}
	if hdrForwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		return strings.Join(parts, ",")
	}
	return hdrRealIP
}
