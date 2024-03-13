package log

import "log/slog"

func Start() {
	slog.Info(`                                                                 
	■■■■■■■               ■                         ■      ■    ■■■■ 
	■■■■■■■               ■                        ■■■     ■■  ■■■■■ 
	   ■                  ■                        ■■■    ■■■  ■■    
	   ■     ■■■■    ■■■■ ■■■■■■   ■■■■   ■ ■■     ■ ■    ■ ■  ■■    
	   ■    ■   ■■  ■■    ■■   ■  ■   ■■  ■■       ■ ■■  ■■ ■  ■■■■■ 
	   ■    ■   ■■  ■     ■    ■  ■   ■■  ■        ■  ■  ■  ■  ■■■■■ 
	   ■    ■■■■■■  ■     ■    ■  ■■■■■■  ■        ■  ■  ■  ■  ■■    
	   ■    ■       ■     ■    ■  ■       ■        ■  ■■■   ■  ■■    
	   ■    ■■      ■■    ■    ■  ■■      ■        ■   ■■   ■  ■■    
	   ■     ■■■■    ■■■■ ■    ■   ■■■■   ■        ■   ■■   ■  ■■■■■■`)
	slog.Info("Starting server on :8080")
}
