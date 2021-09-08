import { Component, Input, OnInit } from '@angular/core';
import { AbstractControl, FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataService } from '../../_services/data.service';
import { FormUtils } from '../../_utils/form.utils';
import { UniformSubscription } from '../../_models/uniform-subscription';
import { UniformSubscriptionFilter } from '../../../../shared/interfaces/uniform-subscription';
import { AppUtils } from '../../_utils/app.utils';

@Component({
  selector: 'ktb-webhook-settings',
  templateUrl: './ktb-webhook-settings.component.html',
  styleUrls: ['./ktb-webhook-settings.component.scss'],
})
export class KtbWebhookSettingsComponent implements OnInit {

  public _projectName?: string;
  public _subscription?: UniformSubscription;
  private _prevFilter?: UniformSubscriptionFilter;
  public webhookConfigForm = this.formBuilder.group({
    method: ['', [Validators.required]],
    url: ['', [Validators.required, Validators.pattern(FormUtils.URL_PATTERN)]],
    payload: ['', [Validators.required]],
    header: this.formBuilder.array([]),
    proxy: ['', [Validators.pattern(FormUtils.URL_PATTERN)]],
  });

  public webhookMethods = ['POST', 'PUT'];

  public loading = false;

  @Input() subscriptionExists = false;

  @Input()
  get projectName(): string {
    return this._projectName || '';
  }

  set projectName(projectName: string) {
    if (this._projectName !== projectName) {
      this._projectName = projectName;
    }
  }

  @Input()
  get subscription(): UniformSubscription | undefined {
    return this._subscription;
  }

  set subscription(value: UniformSubscription | undefined) {
    if (this._subscription !== value) {
      this._subscription = value;
      this.prevFilter = this.subscription?.filter;
    }
  }

  get prevFilter(): UniformSubscriptionFilter | undefined {
    return this._prevFilter;
  }

  set prevFilter(value: UniformSubscriptionFilter | undefined) {
    if (this._prevFilter !== value) {
      this._prevFilter = AppUtils.copyObject(value);
    }
  }

  get header(): FormArray | null {
    return this.webhookConfigForm.get('header') as FormArray;
  }

  get headerControls(): FormGroup[] {
    return this.header?.controls as FormGroup[] || [];
  }

  constructor(private dataService: DataService, private formBuilder: FormBuilder) {
  }

  ngOnInit(): void {
    if (this.subscriptionExists) {
      this.loading = true;
      const stage: string | undefined = this.subscription?.filter?.stages?.length ? this.subscription?.filter?.stages[0] : undefined;
      const services: string | undefined = this.subscription?.filter?.services?.length ? this.subscription?.filter?.services[0] : undefined;
      this.dataService.getWebhookConfig(this.projectName, stage, services)
        .subscribe(webhookConfig => {
          this.webhookConfigForm?.get('method')?.setValue(webhookConfig.method);
          this.webhookConfigForm?.get('url')?.setValue(webhookConfig.url);
          this.webhookConfigForm?.get('payload')?.setValue(webhookConfig.payload);
          this.webhookConfigForm?.get('proxy')?.setValue(webhookConfig.proxy);

          for (const header of webhookConfig.header || []) {
            this.addHeader(header.name, header.value);
          }

          this.loading = false;
        }, () => {
          this.loading = false;
        });
    }
  }

  public addHeader(name?: string, value?: string): void {
    this.header?.push(this.formBuilder.group({
      name: [name, [Validators.required]],
      value: [value, [Validators.required]],
    }));
  }

  public removeHeader(index: number): void {
    this.header?.removeAt(index);
  }

  public getFormControl(controlName: string): AbstractControl | null {
    return this.webhookConfigForm.get(controlName);
  }

}