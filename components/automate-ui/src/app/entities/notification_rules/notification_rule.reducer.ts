import { EntityState, EntityAdapter, createEntityAdapter } from '@ngrx/entity';
import { HttpErrorResponse } from '@angular/common/http';
import { set, pipe, unset } from 'lodash/fp';
import { EntityStatus } from 'app/entities/entities';
import { NotificationRuleActionTypes, NotificationRuleActions } from './notification_rule.action';
import { NotificationRule } from './notification_rule.model';

export interface NotificationRuleEntityState extends EntityState<NotificationRule> {
  rulesStatus:  EntityStatus;
  getAllStatus: EntityStatus;
  saveStatus:   EntityStatus;
  saveError:    HttpErrorResponse;
  deleteStatus: EntityStatus;
}

const GET_ALL_STATUS = 'getAllStatus';
const SAVE_STATUS = 'saveStatus';
const SAVE_ERROR = 'saveError';
const DELETE_STATUS = 'deleteStatus';

export const notificationRuleEntityAdapter:
  EntityAdapter<NotificationRule> = createEntityAdapter<NotificationRule>({
  selectId: (rule: NotificationRule) => rule.name
});

export const NotificationRuleEntityInitialState: NotificationRuleEntityState =
  notificationRuleEntityAdapter.getInitialState(<NotificationRuleEntityState>{
    getAllStatus: EntityStatus.notLoaded,
    deleteStatus: EntityStatus.notLoaded,
  });

export function notificationRuleEntityReducer(
  state: NotificationRuleEntityState = NotificationRuleEntityInitialState,
  action: NotificationRuleActions): NotificationRuleEntityState {

  switch (action.type) {
    case NotificationRuleActionTypes.GET_ALL:
      return set(GET_ALL_STATUS, EntityStatus.loading,
        notificationRuleEntityAdapter.removeAll(state));

    case NotificationRuleActionTypes.GET_ALL_SUCCESS:
      return pipe(
        set(GET_ALL_STATUS, EntityStatus.loadingSuccess))
        (notificationRuleEntityAdapter.setAll(action.payload, state));

    case NotificationRuleActionTypes.GET_ALL_FAILURE:
      return set(GET_ALL_STATUS, EntityStatus.loadingFailure, state);

    case NotificationRuleActionTypes.CREATE: {
      return set(SAVE_STATUS, EntityStatus.loading, state) as NotificationRuleEntityState;
    }

    case NotificationRuleActionTypes.CREATE_SUCCESS: {
      return pipe(
        unset(SAVE_ERROR),
        set(SAVE_STATUS, EntityStatus.loadingSuccess)
      )(notificationRuleEntityAdapter.addOne(action.payload, state)) as NotificationRuleEntityState;
    }

    case NotificationRuleActionTypes.CREATE_FAILURE: {
      return pipe(
        set(SAVE_ERROR, action.payload),
        set(SAVE_STATUS, EntityStatus.loadingFailure)
      )(state) as NotificationRuleEntityState;
    }
    
    case NotificationRuleActionTypes.DELETE:
      return set(DELETE_STATUS, EntityStatus.loading, state);

    case NotificationRuleActionTypes.DELETE_SUCCESS:
      return set(DELETE_STATUS, EntityStatus.loadingSuccess,
        notificationRuleEntityAdapter.removeOne(action.payload.rule.name, state));

    case NotificationRuleActionTypes.DELETE_FAILURE:
      return set(DELETE_STATUS, EntityStatus.loadingFailure, state);

    default:
      return state;
  }
}
