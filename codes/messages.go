package codes

var messages = map[string]string{
	"":                                               "",
	"tx_success":                                     "Success",
	"tx_failed":                                      "Transaction failed",
	"tx_too_early":                                   "Too early",
	"tx_too_late":                                    "Too late",
	"tx_missing_operation":                           "Missing Operation",
	"tx_bad_auth":                                    "Bad auth",
	"tx_no_account":                                  "No source account",
	"tx_bad_auth_extra":                              "Bad auth extra",
	"tx_internal_error":                              "Internal error",
	"tx_account_blocked":                             "Account blocked",
	"tx_duplication":                                 "Transaction duplication",
	"op_inner":                                       "Op inner",
	"op_bad_auth":                                    "You dont have permission to complete this action",
	"op_no_account":                                  "Target account does not exist",
	"op_account_blocked":                             "Operations from blocked account are not allowed",
	"op_no_counterparty":                             "No counterparty",
	"op_counterparty_blocked":                        "Counterparty account is blocked",
	"op_counterparty_wrong_type":                     "Counterparty has wrong account type",
	"op_bad_auth_extra":                              "Bad auth extra",
	"op_account_type_mismatched":                     "Wrong account type in operation. Refresh page and try again",
	"op_type_not_allowed":                            "Type of account you trying to create is not allowed",
	"op_name_duplication":                            "Name duplication",
	"op_referrer_not_found":                          "Referrer not found",
	"op_invalid_account_version":                     "Invalid package version",
	"op_invoice_not_found":                           "Invoice not found",
	"op_invoice_wrong_amount":                        "Amount must be a positive number",
	"op_invoice_balance_mismatch":                    "This account id has no such balance",
	"op_invoice_account_mismatch":                    "This account id has no such balance",
	"op_invoice_already_paid":                        "Invoice have already been paid",
	"op_too_many_signers":                            "Signers limit is exceeded",
	"op_threshold_out_of_range":                      "Threshold out of range",
	"op_bad_signer":                                  "Invalid signer",
	"op_trust_malformed":                             "Trust malformed",
	"op_trust_too_many":                              "Trust many",
	"op_invalid_signer_version":                      "Invalid package version",
	"op_invalid_fee_type":                            "Invalid fee type",
	"op_malformed_range":                             "Invalid range",
	"op_range_overlap":                               "Range you entered overlapped with another one. Delete or reduce an old one before creating new",
	"op_sub_type_not_exist":                          "Subtype not exist",
	"op_not_allowed":                                 "Not allowed",
	"op_type_mismatch":                               "Type mismatch",
	"op_invalid_amount":                              "Invalid amount",
	"op_balance_mismatch":                            "Token asset of balance and token asset of operation are not equal",
	"op_reviewer_not_found":                          "Reviewer not found",
	"op_invalid_details":                             "Invalid details",
	"op_fee_mismatch":                                "Fees mismatched",
	"op_old_signer_not_found":                        "Old signer not found",
	"op_signer_already_exists":                       "Signer already exist",
	"op_destination_not_found":                       "Destination not found",
	"op_request_not_found":                           "Request not found",
	"op_asset_already_exists":                        "Token already exists",
	"op_invalid_max_issuance_amount":                 "Invalid max issuance amount",
	"op_invalid_code":                                "Invalid token code",
	"op_invalid_name":                                "Invalid token name",
	"op_request_already_exists":                      "This request already exists in the system",
	"op_stats_asset_already_exists":                  "It can be only one stats asset in the system",
	"op_line_full":                                   "Payment would cause a destination account to exceed their declared trust limit for the token being sent",
	"op_fee_mismatched":                              "Fees mismatched",
	"op_balance_account_mismatched":                  "Account id has no such balance",
	"op_balance_assets_mismatched":                   "Token asset of balance and token asset of operation are not equal",
	"op_src_balance_not_found":                       "Source balance not found",
	"op_reference_duplication":                       "You cannot make two issuances with the same reference",
	"op_stats_overflow":                              "Overflow during statistics calculation",
	"op_limits_exceeded":                             "Limits exceeded",
	"op_not_allowed_by_asset_policy":                 "This action is not allowed by token policy",
	"op_no_trust":                                    "No trust",
	"op_already_exists":                              "Entry already exists",
	"op_invalid_asset":                               "Invalid token asset",
	"op_invalid_action":                              "Invalid action",
	"op_invalid_policies":                            "Invalid policies",
	"op_asset_not_found":                             "Token not found",
	"op_pair_not_traded":                             "Token not tradable",
	"op_underfunded":                                 "Not enough funds. Reduce the amount and try again",
	"op_cross_self":                                  "Current order crosses your existing order",
	"op_offer_overflow":                              "Failed to create offer",
	"op_asset_pair_not_tradable":                     "Token not tradable",
	"op_physical_price_restriction":                  "Price cannot be lower than physical",
	"op_current_price_restriction":                   "Price cannot be lower than current",
	"op_invalid_percent_fee":                         "Invalid percent fee",
	"op_insufficient_price":                          "Order insufficient price",
	"op_success":                                     "Success",
	"op_malformed":                                   "Operation you are trying to create is malformed in some way",
	"op_balance_not_found":                           "Balance not found",
	"op_invoice_overflow":                            "Failed to create invoice",
	"op_not_found":                                   "Not found",
	"op_too_many_invoices":                           "Too many invoices",
	"op_can_not_delete_in_progress":                  "Cannot delete request while it is progress",
	"op_invalid_external_details":                    "External details are too long",
	"op_asset_is_not_withdrawable":                   "It is not allowed to withdraw specified asset",
	"op_conversion_price_is_not_available":           "Conversion price is not available",
	"op_conversion_overflow":                         "Overflow during conversion",
	"op_converted_amount_mismatched":                 "Specified converted amount does not match calculated",
	"op_balance_lock_overflow":                       "Too much assets are locked in specified balance",
	"op_invalid_universal_amount":                    "Unexpected universal amount value",
	"op_initial_preissued_exceeds_max_issuance":      "Number of tokens available for issuance exceeds max number of tokens to be issued",
	"op_base_asset_or_asset_request_not_found":       "Asset of asset creation request for base asset not found",
	"op_quote_asset_not_found":                       "Quote asset not found",
	"op_start_end_invalid":                           "IO should not end before start",
	"op_invalid_end":                                 "Trying to create IO which already ended",
	"op_invalid_price":                               "Price can not be 0",
	"op_invalid_cap":                                 "Soft cap should not exceed Hard cap",
	"op_insufficient_max_issuance":                   "Max number of tokens can be issued is not sufficient to fulfill soft cap",
	"op_invalid_asset_pair":                          "One of the assets (base or quote) has invalid code or they are equal",
	"op_request_or_sale_already_exists":              "IO creation request or IO already exists for specified token",
	"op_not_authorized":                              "Account not authorized to perform issuance of the asset",
	"op_exceeds_max_issuance_amount":                 "Maximal issuance amount will be exceeded after issuance",
	"op_receiver_full_line":                          "Total funds of receiver will exceed UINT64_MAX after issuance",
	"op_fee_exceeds_amount":                          "Fee is more than amount to issue",
	"op_order_book_does_not_exists":                  "Specified IO does not exists or already closed",
	"op_sale_is_not_started_yet":                     "IO has not been started yet",
	"op_sale_already_ended":                          "IO already ended",
	"op_order_violates_hard_cap":                     "Offer violates hard cap restriction of the IO",
	"op_cant_participate_own_sale":                   "Can not participate in the own IO",
	"op_asset_mismatched":                            "Assets mismatched",
	"op_price_does_not_match":                        "Prices does not match",
	"op_insufficient_preissued":                      "Insufficient amount of tokens available for issuance",
	"op_limits_update_request_reference_duplication": "Reference duplication found while creating limits update request",
}

func getMessage(rawCode string) string {
	return messages[rawCode]
}
