package relations

import (
	"database/sql"
	"policyAuth/internal/models/relations"

	"github.com/gofiber/fiber/v2"
)

// ResourcePermissionHandler handles resource-permission-related requests
type ResourcePermissionHandler struct {
	DB *sql.DB
}

func (h *ResourcePermissionHandler) GetResourcePermissions(c *fiber.Ctx) error {
	rows, err := h.DB.Query("SELECT resource_id, permission_id FROM pds_resource_permission")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer rows.Close()

	var resourcePermissions []relations.ResourcePermission
	for rows.Next() {
		var resourcePermission relations.ResourcePermission
		if err := rows.Scan(&resourcePermission.ResourceID, &resourcePermission.PermissionID); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		resourcePermissions = append(resourcePermissions, resourcePermission)
	}

	return c.JSON(resourcePermissions)
}

func (h *ResourcePermissionHandler) CreateResourcePermission(c *fiber.Ctx) error {
	var resourcePermission relations.ResourcePermission
	if err := c.BodyParser(&resourcePermission); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	_, err := h.DB.Exec("INSERT INTO pds_resource_permission (resource_id, permission_id) VALUES ($1, $2)", resourcePermission.ResourceID, resourcePermission.PermissionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(resourcePermission)
}

func (h *ResourcePermissionHandler) DeleteResourcePermission(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := h.DB.Exec("DELETE FROM pds_resource_permission WHERE id=$1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

