// Package identity
/*
Copyright © 2023 grarich <grarich@grawlily.com>
*/
package identity

import (
	"conoha_cli/internal/apps"
	"conoha_cli/internal/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	endpoint string
	tenantID string
	Cmd      = &cobra.Command{
		Use:   "identity",
		Short: "Identity API を操作します。",
		Long: `Identityの機能を使用できます。
バージョン情報取得、トークン発行が可能です。`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if endpoint != "" && !utils.IsValidUrl(endpoint) {
				return fmt.Errorf("不正なURLです。 URL: %+v", endpoint)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if endpoint != "" {
				if err := apps.SetConfig("endpoint.identity", endpoint); err != nil {
					return err
				}
				cmd.Println("エンドポイントを設定しました。")
			}
			if tenantID != "" {
				if err := apps.SetConfig("identity.tenant_id", tenantID); err != nil {
					return err
				}
				cmd.Println("テナントを設定しました。")
			}
			return nil
		},
	}
)

func init() {
	Cmd.PersistentFlags().StringVar(
		&endpoint,
		"set-endpoint",
		"",
		`identityのAPIを使用する前にエンドポイントの設定をする必要があります。
ConoHaのダッシュボードにログインし、エンドポイントを確認してください。`)
	Cmd.Flags().StringVar(
		&tenantID,
		"set-tenant-id",
		"",
		`テナントIDを設定します。
ConoHaのダッシュボードにログインし、テナントIDを確認してください。`)
	Cmd.AddCommand(GetTokenCmd)
}
