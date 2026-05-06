package bootstrap

import "testing"

func TestBuildServer(t *testing.T) {
	t.Parallel()

	server, err := BuildServer()
	if err != nil {
		t.Fatalf("BuildServer() error = %v", err)
	}
	if server == nil {
		t.Fatal("BuildServer() returned nil")
	}
	if server.App == nil {
		t.Fatal("BuildServer().App returned nil")
	}
	if server.HumaAPI == nil {
		t.Fatal("BuildServer().HumaAPI returned nil")
	}
}

func TestBuildServer_CreatesIsolatedApps(t *testing.T) {
	t.Parallel()

	first, err := BuildServer()
	if err != nil {
		t.Fatalf("BuildServer() first error = %v", err)
	}
	second, err := BuildServer()
	if err != nil {
		t.Fatalf("BuildServer() second error = %v", err)
	}
	if first.App == second.App {
		t.Fatal("BuildServer() reused Fiber app instance")
	}
}

func TestRegisterRoutes(t *testing.T) {
	t.Parallel()

	server, err := BuildServer()
	if err != nil {
		t.Fatalf("BuildServer() error = %v", err)
	}
	if err := RegisterRoutes(server); err != nil {
		t.Fatalf("RegisterRoutes() error = %v", err)
	}
	if len(server.App.GetRoutes()) == 0 {
		t.Fatal("RegisterRoutes() did not register any route")
	}
}
