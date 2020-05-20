
import { createSelector, createFeatureSelector } from '@ngrx/store';

import { NotificationRuleEntityState, notificationRuleEntityAdapter } from './notification_rule.reducer';

export const notificationRuleState = createFeatureSelector<NotificationRuleEntityState>('notificationRules');

export const {
  selectAll: allRules,
  selectEntities: ruleEntities
} = notificationRuleEntityAdapter.getSelectors(notificationRuleState);

export const notificationRuleStatus = createSelector(
  notificationRuleState,
  (state) => state.rulesStatus
);

export const getAllStatus = createSelector(
  notificationRuleState,
  (state) => state.getAllStatus
);

export const saveStatus = createSelector(
  notificationRuleState,
  (state) => state.saveStatus
);

export const saveError = createSelector(
  notificationRuleState,
  (state) => state.saveError
);

export const deleteStatus = createSelector(
  notificationRuleState,
  (state) => state.deleteStatus
);
