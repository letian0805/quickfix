//Package crossordercancelreplacerequest msg type = t.
package crossordercancelreplacerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a CrossOrderCancelReplaceRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//OrderID is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from CrossOrderCancelReplaceRequest.
func (m Message) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossID is a required field for CrossOrderCancelReplaceRequest.
func (m Message) CrossID() (*field.CrossIDField, errors.MessageRejectError) {
	f := &field.CrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossID reads a CrossID from CrossOrderCancelReplaceRequest.
func (m Message) GetCrossID(f *field.CrossIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigCrossID is a required field for CrossOrderCancelReplaceRequest.
func (m Message) OrigCrossID() (*field.OrigCrossIDField, errors.MessageRejectError) {
	f := &field.OrigCrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigCrossID reads a OrigCrossID from CrossOrderCancelReplaceRequest.
func (m Message) GetOrigCrossID(f *field.OrigCrossIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossType is a required field for CrossOrderCancelReplaceRequest.
func (m Message) CrossType() (*field.CrossTypeField, errors.MessageRejectError) {
	f := &field.CrossTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossType reads a CrossType from CrossOrderCancelReplaceRequest.
func (m Message) GetCrossType(f *field.CrossTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossPrioritization is a required field for CrossOrderCancelReplaceRequest.
func (m Message) CrossPrioritization() (*field.CrossPrioritizationField, errors.MessageRejectError) {
	f := &field.CrossPrioritizationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossPrioritization reads a CrossPrioritization from CrossOrderCancelReplaceRequest.
func (m Message) GetCrossPrioritization(f *field.CrossPrioritizationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSides is a required field for CrossOrderCancelReplaceRequest.
func (m Message) NoSides() (*field.NoSidesField, errors.MessageRejectError) {
	f := &field.NoSidesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSides reads a NoSides from CrossOrderCancelReplaceRequest.
func (m Message) GetNoSides(f *field.NoSidesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from CrossOrderCancelReplaceRequest.
func (m Message) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from CrossOrderCancelReplaceRequest.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from CrossOrderCancelReplaceRequest.
func (m Message) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from CrossOrderCancelReplaceRequest.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from CrossOrderCancelReplaceRequest.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from CrossOrderCancelReplaceRequest.
func (m Message) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from CrossOrderCancelReplaceRequest.
func (m Message) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from CrossOrderCancelReplaceRequest.
func (m Message) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from CrossOrderCancelReplaceRequest.
func (m Message) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from CrossOrderCancelReplaceRequest.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from CrossOrderCancelReplaceRequest.
func (m Message) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from CrossOrderCancelReplaceRequest.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from CrossOrderCancelReplaceRequest.
func (m Message) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from CrossOrderCancelReplaceRequest.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from CrossOrderCancelReplaceRequest.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from CrossOrderCancelReplaceRequest.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from CrossOrderCancelReplaceRequest.
func (m Message) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from CrossOrderCancelReplaceRequest.
func (m Message) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from CrossOrderCancelReplaceRequest.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from CrossOrderCancelReplaceRequest.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from CrossOrderCancelReplaceRequest.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from CrossOrderCancelReplaceRequest.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from CrossOrderCancelReplaceRequest.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from CrossOrderCancelReplaceRequest.
func (m Message) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from CrossOrderCancelReplaceRequest.
func (m Message) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from CrossOrderCancelReplaceRequest.
func (m Message) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from CrossOrderCancelReplaceRequest.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from CrossOrderCancelReplaceRequest.
func (m Message) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from CrossOrderCancelReplaceRequest.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from CrossOrderCancelReplaceRequest.
func (m Message) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from CrossOrderCancelReplaceRequest.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from CrossOrderCancelReplaceRequest.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from CrossOrderCancelReplaceRequest.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from CrossOrderCancelReplaceRequest.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from CrossOrderCancelReplaceRequest.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from CrossOrderCancelReplaceRequest.
func (m Message) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from CrossOrderCancelReplaceRequest.
func (m Message) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from CrossOrderCancelReplaceRequest.
func (m Message) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from CrossOrderCancelReplaceRequest.
func (m Message) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from CrossOrderCancelReplaceRequest.
func (m Message) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from CrossOrderCancelReplaceRequest.
func (m Message) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from CrossOrderCancelReplaceRequest.
func (m Message) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from CrossOrderCancelReplaceRequest.
func (m Message) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from CrossOrderCancelReplaceRequest.
func (m Message) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) SettlType() (*field.SettlTypeField, errors.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from CrossOrderCancelReplaceRequest.
func (m Message) GetSettlType(f *field.SettlTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) SettlDate() (*field.SettlDateField, errors.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from CrossOrderCancelReplaceRequest.
func (m Message) GetSettlDate(f *field.SettlDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from CrossOrderCancelReplaceRequest.
func (m Message) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from CrossOrderCancelReplaceRequest.
func (m Message) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from CrossOrderCancelReplaceRequest.
func (m Message) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from CrossOrderCancelReplaceRequest.
func (m Message) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) ExDestination() (*field.ExDestinationField, errors.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from CrossOrderCancelReplaceRequest.
func (m Message) GetExDestination(f *field.ExDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from CrossOrderCancelReplaceRequest.
func (m Message) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) ProcessCode() (*field.ProcessCodeField, errors.MessageRejectError) {
	f := &field.ProcessCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from CrossOrderCancelReplaceRequest.
func (m Message) GetProcessCode(f *field.ProcessCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) PrevClosePx() (*field.PrevClosePxField, errors.MessageRejectError) {
	f := &field.PrevClosePxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from CrossOrderCancelReplaceRequest.
func (m Message) GetPrevClosePx(f *field.PrevClosePxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) LocateReqd() (*field.LocateReqdField, errors.MessageRejectError) {
	f := &field.LocateReqdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from CrossOrderCancelReplaceRequest.
func (m Message) GetLocateReqd(f *field.LocateReqdField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for CrossOrderCancelReplaceRequest.
func (m Message) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from CrossOrderCancelReplaceRequest.
func (m Message) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) NoStipulations() (*field.NoStipulationsField, errors.MessageRejectError) {
	f := &field.NoStipulationsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from CrossOrderCancelReplaceRequest.
func (m Message) GetNoStipulations(f *field.NoStipulationsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for CrossOrderCancelReplaceRequest.
func (m Message) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from CrossOrderCancelReplaceRequest.
func (m Message) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from CrossOrderCancelReplaceRequest.
func (m Message) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from CrossOrderCancelReplaceRequest.
func (m Message) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from CrossOrderCancelReplaceRequest.
func (m Message) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) Spread() (*field.SpreadField, errors.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from CrossOrderCancelReplaceRequest.
func (m Message) GetSpread(f *field.SpreadField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from CrossOrderCancelReplaceRequest.
func (m Message) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) BenchmarkCurveName() (*field.BenchmarkCurveNameField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from CrossOrderCancelReplaceRequest.
func (m Message) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, errors.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from CrossOrderCancelReplaceRequest.
func (m Message) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) BenchmarkPrice() (*field.BenchmarkPriceField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from CrossOrderCancelReplaceRequest.
func (m Message) GetBenchmarkPrice(f *field.BenchmarkPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from CrossOrderCancelReplaceRequest.
func (m Message) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from CrossOrderCancelReplaceRequest.
func (m Message) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from CrossOrderCancelReplaceRequest.
func (m Message) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) YieldType() (*field.YieldTypeField, errors.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from CrossOrderCancelReplaceRequest.
func (m Message) GetYieldType(f *field.YieldTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) Yield() (*field.YieldField, errors.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from CrossOrderCancelReplaceRequest.
func (m Message) GetYield(f *field.YieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) YieldCalcDate() (*field.YieldCalcDateField, errors.MessageRejectError) {
	f := &field.YieldCalcDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from CrossOrderCancelReplaceRequest.
func (m Message) GetYieldCalcDate(f *field.YieldCalcDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) YieldRedemptionDate() (*field.YieldRedemptionDateField, errors.MessageRejectError) {
	f := &field.YieldRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from CrossOrderCancelReplaceRequest.
func (m Message) GetYieldRedemptionDate(f *field.YieldRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) YieldRedemptionPrice() (*field.YieldRedemptionPriceField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from CrossOrderCancelReplaceRequest.
func (m Message) GetYieldRedemptionPrice(f *field.YieldRedemptionPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) YieldRedemptionPriceType() (*field.YieldRedemptionPriceTypeField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from CrossOrderCancelReplaceRequest.
func (m Message) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from CrossOrderCancelReplaceRequest.
func (m Message) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) ComplianceID() (*field.ComplianceIDField, errors.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from CrossOrderCancelReplaceRequest.
func (m Message) GetComplianceID(f *field.ComplianceIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIID is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) IOIID() (*field.IOIIDField, errors.MessageRejectError) {
	f := &field.IOIIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIID reads a IOIID from CrossOrderCancelReplaceRequest.
func (m Message) GetIOIID(f *field.IOIIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from CrossOrderCancelReplaceRequest.
func (m Message) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from CrossOrderCancelReplaceRequest.
func (m Message) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from CrossOrderCancelReplaceRequest.
func (m Message) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) ExpireDate() (*field.ExpireDateField, errors.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from CrossOrderCancelReplaceRequest.
func (m Message) GetExpireDate(f *field.ExpireDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from CrossOrderCancelReplaceRequest.
func (m Message) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) GTBookingInst() (*field.GTBookingInstField, errors.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from CrossOrderCancelReplaceRequest.
func (m Message) GetGTBookingInst(f *field.GTBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) MaxShow() (*field.MaxShowField, errors.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from CrossOrderCancelReplaceRequest.
func (m Message) GetMaxShow(f *field.MaxShowField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) PegOffsetValue() (*field.PegOffsetValueField, errors.MessageRejectError) {
	f := &field.PegOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from CrossOrderCancelReplaceRequest.
func (m Message) GetPegOffsetValue(f *field.PegOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) PegMoveType() (*field.PegMoveTypeField, errors.MessageRejectError) {
	f := &field.PegMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from CrossOrderCancelReplaceRequest.
func (m Message) GetPegMoveType(f *field.PegMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) PegOffsetType() (*field.PegOffsetTypeField, errors.MessageRejectError) {
	f := &field.PegOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from CrossOrderCancelReplaceRequest.
func (m Message) GetPegOffsetType(f *field.PegOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) PegLimitType() (*field.PegLimitTypeField, errors.MessageRejectError) {
	f := &field.PegLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from CrossOrderCancelReplaceRequest.
func (m Message) GetPegLimitType(f *field.PegLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) PegRoundDirection() (*field.PegRoundDirectionField, errors.MessageRejectError) {
	f := &field.PegRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from CrossOrderCancelReplaceRequest.
func (m Message) GetPegRoundDirection(f *field.PegRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) PegScope() (*field.PegScopeField, errors.MessageRejectError) {
	f := &field.PegScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from CrossOrderCancelReplaceRequest.
func (m Message) GetPegScope(f *field.PegScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) DiscretionInst() (*field.DiscretionInstField, errors.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from CrossOrderCancelReplaceRequest.
func (m Message) GetDiscretionInst(f *field.DiscretionInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) DiscretionOffsetValue() (*field.DiscretionOffsetValueField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from CrossOrderCancelReplaceRequest.
func (m Message) GetDiscretionOffsetValue(f *field.DiscretionOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) DiscretionMoveType() (*field.DiscretionMoveTypeField, errors.MessageRejectError) {
	f := &field.DiscretionMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from CrossOrderCancelReplaceRequest.
func (m Message) GetDiscretionMoveType(f *field.DiscretionMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) DiscretionOffsetType() (*field.DiscretionOffsetTypeField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from CrossOrderCancelReplaceRequest.
func (m Message) GetDiscretionOffsetType(f *field.DiscretionOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) DiscretionLimitType() (*field.DiscretionLimitTypeField, errors.MessageRejectError) {
	f := &field.DiscretionLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from CrossOrderCancelReplaceRequest.
func (m Message) GetDiscretionLimitType(f *field.DiscretionLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) DiscretionRoundDirection() (*field.DiscretionRoundDirectionField, errors.MessageRejectError) {
	f := &field.DiscretionRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from CrossOrderCancelReplaceRequest.
func (m Message) GetDiscretionRoundDirection(f *field.DiscretionRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) DiscretionScope() (*field.DiscretionScopeField, errors.MessageRejectError) {
	f := &field.DiscretionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from CrossOrderCancelReplaceRequest.
func (m Message) GetDiscretionScope(f *field.DiscretionScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) TargetStrategy() (*field.TargetStrategyField, errors.MessageRejectError) {
	f := &field.TargetStrategyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from CrossOrderCancelReplaceRequest.
func (m Message) GetTargetStrategy(f *field.TargetStrategyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) TargetStrategyParameters() (*field.TargetStrategyParametersField, errors.MessageRejectError) {
	f := &field.TargetStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from CrossOrderCancelReplaceRequest.
func (m Message) GetTargetStrategyParameters(f *field.TargetStrategyParametersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) ParticipationRate() (*field.ParticipationRateField, errors.MessageRejectError) {
	f := &field.ParticipationRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from CrossOrderCancelReplaceRequest.
func (m Message) GetParticipationRate(f *field.ParticipationRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) CancellationRights() (*field.CancellationRightsField, errors.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from CrossOrderCancelReplaceRequest.
func (m Message) GetCancellationRights(f *field.CancellationRightsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, errors.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from CrossOrderCancelReplaceRequest.
func (m Message) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from CrossOrderCancelReplaceRequest.
func (m Message) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for CrossOrderCancelReplaceRequest.
func (m Message) Designation() (*field.DesignationField, errors.MessageRejectError) {
	f := &field.DesignationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from CrossOrderCancelReplaceRequest.
func (m Message) GetDesignation(f *field.DesignationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds CrossOrderCancelReplaceRequest messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for CrossOrderCancelReplaceRequest.
func Builder(
	crossid *field.CrossIDField,
	origcrossid *field.OrigCrossIDField,
	crosstype *field.CrossTypeField,
	crossprioritization *field.CrossPrioritizationField,
	nosides *field.NoSidesField,
	transacttime *field.TransactTimeField,
	ordtype *field.OrdTypeField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("t"))
	builder.Body().Set(crossid)
	builder.Body().Set(origcrossid)
	builder.Body().Set(crosstype)
	builder.Body().Set(crossprioritization)
	builder.Body().Set(nosides)
	builder.Body().Set(transacttime)
	builder.Body().Set(ordtype)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "t", r
}