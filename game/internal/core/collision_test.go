package core

import "testing"

func TestCirclesOverlap(t *testing.T) {
	if !circlesOverlap(0, 0, 5, 8, 0, 5) {
		t.Fatal("expected overlap")
	}
	if circlesOverlap(0, 0, 5, 20, 0, 5) {
		t.Fatal("expected no overlap")
	}
}

func TestCollisionMaskPlayerEnemy(t *testing.T) {
	player := NewCollisionMask(LayerPlayer, LayerEnemy)
	enemy := NewCollisionMask(LayerEnemy, LayerPlayer)
	if !player.CanCollideWith(enemy) || !enemy.CanCollideWith(player) {
		t.Fatal("player and enemy should collide")
	}
}

func TestCheckCollisionNotifiesPlayer(t *testing.T) {
	m := NewCollisionManager()

	playerShape := NewCollisionCircle(6)
	playerMask := NewCollisionMask(LayerPlayer, LayerEnemy)
	player := m.NewCollider("player", playerShape, playerMask)
	player.SetPosition(0, 0)

	enemyShape := NewCollisionCircle(6)
	enemyMask := NewCollisionMask(LayerEnemy, LayerPlayer)
	enemy := m.NewCollider("enemy", enemyShape, enemyMask)
	enemy.SetPosition(8, 0)

	hits := 0
	player.SetOnCollide(func(other *Collider) {
		hits++
	})

	m.CheckCollision()
	if hits != 1 {
		t.Fatalf("hits = %d, want 1", hits)
	}
}
