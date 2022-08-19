package helper

var (
	// InternalServer ...
	InternalServer = "internal_server"
	// JWT ...
	JWT = "jwt"
	// InvalidCredential ...
	InvalidCredential = "invalid_credential"
	// InactiveUser ...
	InactiveUser = "inactive_user"
	// InvalidPassword ...
	InvalidPassword = "invalid_password"
	// InvalidEmail ...
	InvalidEmail = "invalid_email"
	// HeaderNotPresent ...
	HeaderNotPresent = "header_not_present"
	// UnexpectedSigningMethod ...
	UnexpectedSigningMethod = "unexpected_signing_method"
	// UnexpectedClaims ...
	UnexpectedClaims = "unexpected_claims"
	// ExpiredToken ...
	ExpiredToken = "expired_token"
	// Unauthorized ...
	Unauthorized = "unauthoized"
	// InvalidRole ...
	InvalidRole = "invalid_role"

	// RecordAlreadyExist ...
	RecordAlreadyExist = "record_already_exist"
	// NameAlreadyExist ...
	NameAlreadyExist = "name_already_exist"
	// ReferencedRecord ...
	ReferencedRecord = "referenced_record"
	// InvalidParameter ...
	InvalidParameter = "invalid_parameter"

	// FileTooBig ...
	FileTooBig = "file_too_big"
	// InvalidFileType ...
	InvalidFileType = "invalid_file_type"
	// InvalidProfileImage ...
	InvalidProfileImage = "invalid_profile_image"
	// InvalidCity ...
	InvalidCity = "invalid_city"
	// CodeAlredyUsed ...
	CodeAlredyUsed = "code_already_used"
	// OrderByNotAllowed ...
	OrderByNotAllowed = "order_by_not_allowed"

	// InactiveAdmin ...
	InactiveAdmin = "inactive_admin"
	// DuplicateEmail ...
	DuplicateEmail = "duplicate_email"
	// DuplicatePhone ...
	DuplicatePhone = "duplicate_phone"
	// InvalidOldPassword ...
	InvalidOldPassword = "invalid_old_password"
	// InvalidIdentityImage ...
	InvalidIdentityImage = "invalid_identity_image"
	// InvalidSelfieImage ...
	InvalidSelfieImage = "invalid_selfie_image"
	// InvalidStatus ...
	InvalidStatus = "invalid_status"
	// InvalidEmailOrPhone ...
	InvalidEmailOrPhone = "invalid_email_phone"

	// InvalidAreaCoverage ...
	InvalidAreaCoverage = "invalid_area_coverage"
	// InvalidDroppedDateFormat ...
	InvalidDroppedDateFormat = "invalid_dropped_date_format"
	// InvalidDeliveredDateFormat ...
	InvalidDeliveredDateFormat = "invalid_delivered_date_format"
	// InvalidAccess ...
	InvalidAccess = "invalid_access"
	// InvalidPoint ...
	InvalidPoint = "invalid_point"
	// InvalidBirthDate ...
	InvalidBirthDate = "invalid_birth_date"
	// InvalidDay ...
	InvalidDay = "invalid_day"
	// InvalidTime ...
	InvalidTime = "invalid_time"
	// ExpKey ...
	ExpKey = "exp_key"
	// PasswordLength ...
	PasswordLength = "password_length"
	// SendMail ...
	SendMail = "send_mail"
	// InvalidKey
	InvalidKey = "invalid_key"

	// UserNotFound ...
	UserNotFound = "user_not_found"
	// ExpPassword ...
	ExpPassword = "exp_password"
	// Recaptcha ...
	Recaptcha = "recaptcha"
	// InvalidLimit ...
	InvalidLimit = "invalid_limit"
	// InvalidPrice ...
	InvalidPrice = "invalid_price"
	// InvalidClient ...
	InvalidClient = "invalid_client"
	// InvalidPeriod ...
	InvalidPeriod = "invalid_period"
	// DataNotFound ...
	DataNotFound = "data_not_found"

	// InvalidMinOrder ...
	InvalidMinOrder = "invalid_min_order"
	// OderLimit ...
	OderLimit = "order_limit"

	// InvalidTypeOtp ...
	InvalidTypeOtp = "invalid_type_otp"
	// PhoneNotValid ...
	PhoneNotValid = "phone_not_valid"
	// SendSMS ...
	SendSms = "send_sms"
	// WrongOTP ...
	WrongOTP = "Wrong OTP"
	//emailexist
	EmailAlreadyExist = "Email Already Exist"

	ReferralNotFound = "Referral Code Not Exist"

	ReferralReachLimit = "Can't Use this Referral Code.\nReferral Code Have Reach Limit"

	ActivePackageOrderNotExist = "You Have not active package"

	UnApproverPackageOrder = "Your package still not approved. \nPlease contact admin to approve your package"

	UnPaidPackageOrder = "Your have uncleared payment. \nPlease compleate your last payment or contact our admin "

	UnPaidPackageOrderForAbsen = "You can't take attendance before compleate the payment"

	DailyLimitAbsen = "You can't take attendance because you have reach daily limit"

	WeeklyLimitAbsen = "You can't take attendance because you have reach weekly limit"

	AbsentCountReachLimit = "You have take 4 attendance.\nPlease paid your package first."

	//apabila coach asli belum membuat jadwal namun member absen dengan coach subtitusi
	AttendanceScheduleNotExist = "Attendace schedule not exist"

	WrongEmailOrPassword = "Email or password is wrong"
)
