import { HttpErrorResponse } from '@angular/common/http';
import { Action } from '@ngrx/store';
import { NotificationRule } from './notification_rule.model';

export enum NotificationRuleActionTypes {
  GET_ALL = 'NOTIFICATION_RULES::GET_ALL',
  GET_ALL_SUCCESS = 'NOTIFICATION_RULES::GET_ALL::SUCCESS',
  GET_ALL_FAILURE = 'NOTIFICATION_RULES::GET_ALL::FAILURE',
  DELETE          = 'NOTIFICATION_RULES::DELETE',
  DELETE_SUCCESS  = 'NOTIFICATION_RULES::DELETE::SUCCESS',
  DELETE_FAILURE  = 'NOTIFICATION_RULES::DELETE::FAILURE',
}

export interface NotificationRulesSuccessPayload {
  rules: NotificationRule[];
}

export class GetNotificationRules implements Action {
  readonly type = NotificationRuleActionTypes.GET_ALL;
}

export class GetNotificationRulesSuccess implements Action {
  readonly type = NotificationRuleActionTypes.GET_ALL_SUCCESS;
  constructor(public payload: NotificationRule[]) { }
}

export class GetNotificationRulesFailure implements Action {
  readonly type = NotificationRuleActionTypes.GET_ALL_FAILURE;
  constructor(public payload: HttpErrorResponse) { }
}

export class DeleteNotificationRule implements Action {
  readonly type = NotificationRuleActionTypes.DELETE;

  constructor(public payload: { rule: NotificationRule }) { }
}

export class DeleteNotificationRuleSuccess implements Action {
  readonly type = NotificationRuleActionTypes.DELETE_SUCCESS;

  constructor(public payload: { rule: NotificationRule }) { }
}

export class DeleteNotificationRuleFailure implements Action {
  readonly type = NotificationRuleActionTypes.DELETE_FAILURE;

  constructor(public payload: HttpErrorResponse) { }
}

export type NotificationRuleActions =
  | GetNotificationRules
  | GetNotificationRulesSuccess
  | GetNotificationRulesFailure
  | DeleteNotificationRule
  | DeleteNotificationRuleSuccess
  | DeleteNotificationRuleFailure;
